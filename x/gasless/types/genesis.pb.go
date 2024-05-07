// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aib/gasless/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the gasless module's genesis state.
type GenesisState struct {
	Params         Params        `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	TxToGasTankIds []TxGTIDs     `protobuf:"bytes,2,rep,name=tx_to_gas_tank_ids,json=txToGasTankIds,proto3" json:"tx_to_gas_tank_ids"`
	LastGasTankId  uint64        `protobuf:"varint,3,opt,name=last_gas_tank_id,json=lastGasTankId,proto3" json:"last_gas_tank_id,omitempty"`
	GasTanks       []GasTank     `protobuf:"bytes,4,rep,name=gas_tanks,json=gasTanks,proto3" json:"gas_tanks"`
	GasConsumers   []GasConsumer `protobuf:"bytes,5,rep,name=gas_consumers,json=gasConsumers,proto3" json:"gas_consumers"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d441cc41805faaaf, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "aib.gasless.v1beta1.GenesisState")
}

func init() { proto.RegisterFile("aib/gasless/v1beta1/genesis.proto", fileDescriptor_d441cc41805faaaf) }

var fileDescriptor_d441cc41805faaaf = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4e, 0xf2, 0x40,
	0x10, 0xc7, 0x5b, 0xe0, 0x23, 0x9f, 0x05, 0x8c, 0xa9, 0x1e, 0x1a, 0x34, 0xb5, 0x7a, 0x91, 0x8b,
	0xdd, 0x00, 0x27, 0x4f, 0x46, 0x30, 0x21, 0xc4, 0xc4, 0x28, 0xf6, 0xe4, 0xa5, 0x99, 0xc2, 0x66,
	0x6d, 0x68, 0xbb, 0x4d, 0x67, 0x31, 0xf5, 0x2d, 0x7c, 0x2c, 0x8e, 0x1c, 0x3d, 0x19, 0x85, 0x17,
	0x31, 0xb4, 0x0b, 0x7a, 0xa8, 0xde, 0x76, 0x67, 0x7e, 0xf3, 0xfb, 0x4f, 0x32, 0xda, 0x09, 0xf8,
	0x1e, 0x61, 0x80, 0x01, 0x45, 0x24, 0xcf, 0x6d, 0x8f, 0x0a, 0x68, 0x13, 0x46, 0x23, 0x8a, 0x3e,
	0xda, 0x71, 0xc2, 0x05, 0xd7, 0xf7, 0xc1, 0xf7, 0x6c, 0x89, 0xd8, 0x12, 0x69, 0x1e, 0x30, 0xce,
	0x78, 0xd6, 0x27, 0xeb, 0x57, 0x8e, 0x36, 0xad, 0x22, 0x5b, 0x0c, 0x09, 0x84, 0x52, 0xd6, 0x2c,
	0xce, 0x93, 0xf2, 0x0c, 0x39, 0x5d, 0x94, 0xb4, 0xfa, 0x20, 0xdf, 0xe0, 0x41, 0x80, 0xa0, 0xfa,
	0x85, 0x56, 0xcd, 0x1d, 0x86, 0x6a, 0xa9, 0xad, 0x5a, 0xe7, 0xd0, 0x2e, 0xd8, 0xc8, 0xbe, 0xcb,
	0x90, 0x5e, 0x65, 0xfe, 0x7e, 0xac, 0x8c, 0xe4, 0x80, 0x7e, 0xab, 0xe9, 0x22, 0x75, 0x05, 0x77,
	0x19, 0xa0, 0x2b, 0x20, 0x9a, 0xba, 0xfe, 0x04, 0x8d, 0x92, 0x55, 0x6e, 0xd5, 0x3a, 0x47, 0x85,
	0x1a, 0x27, 0x1d, 0x38, 0xc3, 0xeb, 0x8d, 0x67, 0x57, 0xa4, 0x0e, 0x1f, 0x00, 0x3a, 0x10, 0x4d,
	0x87, 0x13, 0xd4, 0xcf, 0xb4, 0xbd, 0x00, 0x50, 0xfc, 0xd4, 0x19, 0x65, 0x4b, 0x6d, 0x55, 0x46,
	0x8d, 0x75, 0x7d, 0x4b, 0xea, 0x97, 0xda, 0xce, 0x86, 0x41, 0xa3, 0xf2, 0x47, 0x9e, 0x1c, 0x91,
	0x79, 0xff, 0x59, 0xfe, 0x45, 0xfd, 0x46, 0x6b, 0xac, 0x05, 0x63, 0x1e, 0xe1, 0x2c, 0xa4, 0x09,
	0x1a, 0xff, 0x32, 0x89, 0xf5, 0x9b, 0xa4, 0x2f, 0x41, 0x29, 0xaa, 0xb3, 0xef, 0x12, 0xf6, 0xee,
	0xe7, 0x9f, 0xa6, 0x32, 0x5f, 0x9a, 0xea, 0x62, 0x69, 0xaa, 0x1f, 0x4b, 0x53, 0x7d, 0x5d, 0x99,
	0xca, 0x62, 0x65, 0x2a, 0x6f, 0x2b, 0x53, 0x79, 0xec, 0x32, 0x5f, 0x3c, 0xcd, 0x3c, 0x7b, 0xcc,
	0x43, 0x72, 0x15, 0x04, 0xc3, 0xa8, 0x47, 0x05, 0xf6, 0x79, 0x48, 0xc0, 0xf7, 0xce, 0x13, 0x1e,
	0x04, 0x10, 0xc7, 0x24, 0xdd, 0x5e, 0x4d, 0xbc, 0xc4, 0x14, 0xbd, 0x6a, 0x76, 0xac, 0xee, 0x57,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x98, 0x6f, 0xb0, 0x56, 0x41, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GasConsumers) > 0 {
		for iNdEx := len(m.GasConsumers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GasConsumers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.GasTanks) > 0 {
		for iNdEx := len(m.GasTanks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GasTanks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.LastGasTankId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastGasTankId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.TxToGasTankIds) > 0 {
		for iNdEx := len(m.TxToGasTankIds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TxToGasTankIds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.TxToGasTankIds) > 0 {
		for _, e := range m.TxToGasTankIds {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.LastGasTankId != 0 {
		n += 1 + sovGenesis(uint64(m.LastGasTankId))
	}
	if len(m.GasTanks) > 0 {
		for _, e := range m.GasTanks {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GasConsumers) > 0 {
		for _, e := range m.GasConsumers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxToGasTankIds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxToGasTankIds = append(m.TxToGasTankIds, TxGTIDs{})
			if err := m.TxToGasTankIds[len(m.TxToGasTankIds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastGasTankId", wireType)
			}
			m.LastGasTankId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastGasTankId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasTanks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasTanks = append(m.GasTanks, GasTank{})
			if err := m.GasTanks[len(m.GasTanks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasConsumers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasConsumers = append(m.GasConsumers, GasConsumer{})
			if err := m.GasConsumers[len(m.GasConsumers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
