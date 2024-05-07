package keeper_test

import (
	"encoding/binary"
	"testing"
	"strings"

	"github.com/stretchr/testify/suite"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	chain "github.com/AllInBetsCom/aib-rollapp/app"
	// utils "github.com/AllInBetsCom/aib-rollapp/types"
	"github.com/AllInBetsCom/aib-rollapp/x/gasless"
	"github.com/AllInBetsCom/aib-rollapp/x/gasless/keeper"
	"github.com/AllInBetsCom/aib-rollapp/x/gasless/types"
	minttypes "github.com/dymensionxyz/dymension-rdk/x/mint/types"
)

type KeeperTestSuite struct {
	suite.Suite

	app       *chain.App
	ctx       sdk.Context
	keeper    keeper.Keeper
	querier   keeper.Querier
	msgServer types.MsgServer
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	s.app = chain.Setup(s.T(), false)
	s.ctx = s.app.BaseApp.NewContext(false, tmproto.Header{})
	s.keeper = s.app.GaslessKeeper
	s.querier = keeper.Querier{Keeper: s.keeper}
	s.msgServer = keeper.NewMsgServerImpl(s.keeper)
}

// Below are just shortcuts to frequently-used functions.
func (s *KeeperTestSuite) getBalances(addr sdk.AccAddress) sdk.Coins {
	return s.app.BankKeeper.GetAllBalances(s.ctx, addr)
}

func (s *KeeperTestSuite) getBalance(addr sdk.AccAddress, denom string) sdk.Coin {
	return s.app.BankKeeper.GetBalance(s.ctx, addr, denom)
}

func (s *KeeperTestSuite) sendCoins(fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) {
	s.T().Helper()
	err := s.app.BankKeeper.SendCoins(s.ctx, fromAddr, toAddr, amt)
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) nextBlock() {
	gasless.EndBlocker(s.ctx, s.keeper)
	gasless.BeginBlocker(s.ctx, s.keeper)
}

// Below are useful helpers to write test code easily.
func (s *KeeperTestSuite) addr(addrNum int) sdk.AccAddress {
	addr := make(sdk.AccAddress, 20)
	binary.PutVarint(addr, int64(addrNum))
	return addr
}

func (s *KeeperTestSuite) fundAddr(addr sdk.AccAddress, amt sdk.Coins) {
	// using mint module to mint new test coins, since gasless module is not allowed to mint coins
	s.T().Helper()
	err := s.app.BankKeeper.MintCoins(s.ctx, minttypes.ModuleName, amt)
	s.Require().NoError(err)
	err = s.app.BankKeeper.SendCoinsFromModuleToAccount(s.ctx, minttypes.ModuleName, addr, amt)
	s.Require().NoError(err)
}

func newInt(i int64) sdkmath.Int {
	return sdkmath.NewInt(i)
}

func newDec(i int64) sdkmath.LegacyDec {
	return sdkmath.LegacyNewDec(i)
}

func coinEq(exp, got sdk.Coin) (bool, string, string, string) {
	return exp.IsEqual(got), "expected:\t%v\ngot:\t\t%v", exp.String(), got.String()
}

func coinsEq(exp, got sdk.Coins) (bool, string, string, string) {
	return exp.IsEqual(got), "expected:\t%v\ngot:\t\t%v", exp.String(), got.String()
}

func intEq(exp, got sdkmath.Int) (bool, string, string, string) {
	return exp.Equal(got), "expected:\t%v\ngot:\t\t%v", exp.String(), got.String()
}

func decEq(exp, got sdkmath.LegacyDec) (bool, string, string, string) {
	return exp.Equal(got), "expected:\t%v\ngot:\t\t%v", exp.String(), got.String()
}

func (s *KeeperTestSuite) CreateNewGasTank(
	provider sdk.AccAddress,
	feeDenom string,
	maxFeeUsagePerTx sdkmath.Int,
	maxTxsCountsPerConsumer uint64,
	maxFeeUsagePerConsumer sdkmath.Int,
	txsAllowed, contractsAllowed []string,
	deposit string,
) types.GasTank {
	parsedDepositCoin := ParseCoin(deposit)
	s.fundAddr(provider, sdk.NewCoins(parsedDepositCoin))

	txsAllowed = types.RemoveDuplicates(txsAllowed)
	contractsAllowed = types.RemoveDuplicates(contractsAllowed)
	tank, err := s.keeper.CreateGasTank(s.ctx, types.NewMsgCreateGasTank(
		provider,
		feeDenom,
		maxFeeUsagePerTx,
		maxTxsCountsPerConsumer,
		maxFeeUsagePerConsumer,
		txsAllowed,
		contractsAllowed,
		parsedDepositCoin,
	))
	s.Require().NoError(err)
	s.Require().IsType(types.GasTank{}, tank)
	s.Require().Equal(feeDenom, tank.FeeDenom)
	s.Require().Equal(maxFeeUsagePerTx, tank.MaxFeeUsagePerTx)
	s.Require().Equal(maxTxsCountsPerConsumer, tank.MaxTxsCountPerConsumer)
	s.Require().Equal(maxFeeUsagePerConsumer, tank.MaxFeeUsagePerConsumer)
	s.Require().Equal(txsAllowed, tank.TxsAllowed)
	s.Require().Equal(contractsAllowed, tank.ContractsAllowed)
	s.Require().Equal(ParseCoin(deposit), s.getBalance(tank.GetGasTankReserveAddress(), feeDenom))

	for _, tx := range txsAllowed {
		txGtids, found := s.keeper.GetTxGTIDs(s.ctx, tx)
		s.Require().True(found)
		s.Require().IsType(types.TxGTIDs{}, txGtids)
		s.Require().IsType([]uint64{}, txGtids.GasTankIds)
		s.Require().Equal(txGtids.TxPathOrContractAddress, tx)
		s.Require().Equal(tank.Id, txGtids.GasTankIds[len(txGtids.GasTankIds)-1])
	}

	for _, c := range contractsAllowed {
		txGtids, found := s.keeper.GetTxGTIDs(s.ctx, c)
		s.Require().True(found)
		s.Require().IsType(types.TxGTIDs{}, txGtids)
		s.Require().IsType([]uint64{}, txGtids.GasTankIds)
		s.Require().Equal(txGtids.TxPathOrContractAddress, c)
		s.Require().Equal(tank.Id, txGtids.GasTankIds[len(txGtids.GasTankIds)-1])
	}

	return tank
}

// ParseCoin parses and returns sdk.Coin.
func ParseCoin(s string) sdk.Coin {
	coin, err := sdk.ParseCoinNormalized(strings.ReplaceAll(s, "_", ""))
	if err != nil {
		panic(err)
	}
	return coin
}