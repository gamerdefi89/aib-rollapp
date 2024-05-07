// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: aib/cron/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgRegisterContract struct {
	SecurityAddress string `protobuf:"bytes,1,opt,name=security_address,json=securityAddress,proto3" json:"security_address,omitempty"`
	GameName        string `protobuf:"bytes,2,opt,name=game_name,json=gameName,proto3" json:"game_name,omitempty"`
	ContractAddress string `protobuf:"bytes,3,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	GameType        uint64 `protobuf:"varint,4,opt,name=game_type,json=gameType,proto3" json:"game_type,omitempty"`
}

func (m *MsgRegisterContract) Reset()         { *m = MsgRegisterContract{} }
func (m *MsgRegisterContract) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterContract) ProtoMessage()    {}
func (*MsgRegisterContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_788f49e6e4335c77, []int{0}
}
func (m *MsgRegisterContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterContract.Merge(m, src)
}
func (m *MsgRegisterContract) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterContract) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterContract.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterContract proto.InternalMessageInfo

func (m *MsgRegisterContract) GetSecurityAddress() string {
	if m != nil {
		return m.SecurityAddress
	}
	return ""
}

func (m *MsgRegisterContract) GetGameName() string {
	if m != nil {
		return m.GameName
	}
	return ""
}

func (m *MsgRegisterContract) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *MsgRegisterContract) GetGameType() uint64 {
	if m != nil {
		return m.GameType
	}
	return 0
}

type MsgRegisterContractResponse struct {
}

func (m *MsgRegisterContractResponse) Reset()         { *m = MsgRegisterContractResponse{} }
func (m *MsgRegisterContractResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterContractResponse) ProtoMessage()    {}
func (*MsgRegisterContractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_788f49e6e4335c77, []int{1}
}
func (m *MsgRegisterContractResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterContractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterContractResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterContractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterContractResponse.Merge(m, src)
}
func (m *MsgRegisterContractResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterContractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterContractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterContractResponse proto.InternalMessageInfo

type MsgDeRegisterContract struct {
	SecurityAddress string `protobuf:"bytes,1,opt,name=security_address,json=securityAddress,proto3" json:"security_address,omitempty"`
	GameId          uint64 `protobuf:"varint,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
}

func (m *MsgDeRegisterContract) Reset()         { *m = MsgDeRegisterContract{} }
func (m *MsgDeRegisterContract) String() string { return proto.CompactTextString(m) }
func (*MsgDeRegisterContract) ProtoMessage()    {}
func (*MsgDeRegisterContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_788f49e6e4335c77, []int{2}
}
func (m *MsgDeRegisterContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeRegisterContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeRegisterContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeRegisterContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeRegisterContract.Merge(m, src)
}
func (m *MsgDeRegisterContract) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeRegisterContract) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeRegisterContract.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeRegisterContract proto.InternalMessageInfo

func (m *MsgDeRegisterContract) GetSecurityAddress() string {
	if m != nil {
		return m.SecurityAddress
	}
	return ""
}

func (m *MsgDeRegisterContract) GetGameId() uint64 {
	if m != nil {
		return m.GameId
	}
	return 0
}

type MsgDeRegisterContractResponse struct {
}

func (m *MsgDeRegisterContractResponse) Reset()         { *m = MsgDeRegisterContractResponse{} }
func (m *MsgDeRegisterContractResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeRegisterContractResponse) ProtoMessage()    {}
func (*MsgDeRegisterContractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_788f49e6e4335c77, []int{3}
}
func (m *MsgDeRegisterContractResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeRegisterContractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeRegisterContractResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeRegisterContractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeRegisterContractResponse.Merge(m, src)
}
func (m *MsgDeRegisterContractResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeRegisterContractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeRegisterContractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeRegisterContractResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRegisterContract)(nil), "aib.cron.v1beta1.MsgRegisterContract")
	proto.RegisterType((*MsgRegisterContractResponse)(nil), "aib.cron.v1beta1.MsgRegisterContractResponse")
	proto.RegisterType((*MsgDeRegisterContract)(nil), "aib.cron.v1beta1.MsgDeRegisterContract")
	proto.RegisterType((*MsgDeRegisterContractResponse)(nil), "aib.cron.v1beta1.MsgDeRegisterContractResponse")
}

func init() { proto.RegisterFile("aib/cron/v1beta1/tx.proto", fileDescriptor_788f49e6e4335c77) }

var fileDescriptor_788f49e6e4335c77 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xb1, 0x6e, 0xe2, 0x30,
	0x1c, 0xc6, 0xf1, 0x81, 0xb8, 0xc3, 0xcb, 0xa1, 0xdc, 0x9d, 0x80, 0x20, 0x72, 0x28, 0xd2, 0xe9,
	0xe8, 0x40, 0x5c, 0xda, 0x27, 0x00, 0xba, 0x50, 0x89, 0x0e, 0x51, 0xa7, 0x76, 0x40, 0x4e, 0x62,
	0x99, 0x48, 0x71, 0x1c, 0xd9, 0x06, 0xc1, 0x5b, 0xf4, 0x25, 0xfa, 0x2e, 0x1d, 0x19, 0xbb, 0xb5,
	0x82, 0x17, 0xa9, 0x12, 0x27, 0x1d, 0x20, 0x03, 0x52, 0xa7, 0xc4, 0xff, 0xef, 0xf3, 0xef, 0xf3,
	0x27, 0x1b, 0x76, 0x70, 0xe8, 0x21, 0x5f, 0xf0, 0x18, 0xad, 0x47, 0x1e, 0x51, 0x78, 0x84, 0xd4,
	0xc6, 0x49, 0x04, 0x57, 0xdc, 0x68, 0xe2, 0xd0, 0x73, 0x52, 0xc9, 0xc9, 0x25, 0xf3, 0x37, 0xe5,
	0x94, 0x67, 0x22, 0x4a, 0xff, 0xb4, 0xcf, 0x6c, 0xf9, 0x5c, 0x32, 0x2e, 0x11, 0x93, 0x14, 0xad,
	0x47, 0xe9, 0x27, 0x17, 0x3a, 0x5a, 0x58, 0xe8, 0x1d, 0x7a, 0x91, 0x4b, 0xbd, 0x93, 0xd8, 0x04,
	0x0b, 0xcc, 0x72, 0xd9, 0x7e, 0x06, 0xf0, 0xd7, 0x5c, 0x52, 0x97, 0xd0, 0x50, 0x2a, 0x22, 0xa6,
	0x3c, 0x56, 0x02, 0xfb, 0xca, 0xb8, 0x80, 0x4d, 0x49, 0xfc, 0x95, 0x08, 0xd5, 0x76, 0x81, 0x83,
	0x40, 0x10, 0x29, 0xdb, 0xa0, 0x0f, 0x06, 0x0d, 0xf7, 0x67, 0x31, 0x1f, 0xeb, 0xb1, 0xd1, 0x85,
	0x0d, 0x8a, 0x19, 0x59, 0xc4, 0x98, 0x91, 0xf6, 0xb7, 0xcc, 0xf3, 0x23, 0x1d, 0xdc, 0x61, 0x46,
	0x52, 0x8e, 0x9f, 0x33, 0x3f, 0x39, 0x55, 0xcd, 0x29, 0xe6, 0xc7, 0x1c, 0xb5, 0x4d, 0x48, 0xbb,
	0xd6, 0x07, 0x83, 0x9a, 0xe6, 0xdc, 0x6f, 0x13, 0x62, 0xf7, 0x60, 0xb7, 0xe4, 0x98, 0x2e, 0x91,
	0x09, 0x8f, 0x25, 0xb1, 0x1f, 0xe1, 0x9f, 0xb9, 0xa4, 0x37, 0xe4, 0x2b, 0x3d, 0x5a, 0xf0, 0x7b,
	0x96, 0x1f, 0x06, 0x59, 0x8b, 0x9a, 0x5b, 0x4f, 0x97, 0xb3, 0xc0, 0xfe, 0x0b, 0x7b, 0xa5, 0xf0,
	0x22, 0xfd, 0xea, 0x0d, 0xc0, 0xea, 0x5c, 0x52, 0x63, 0x09, 0x9b, 0x27, 0x07, 0xf8, 0xe7, 0x1c,
	0x5f, 0xae, 0x53, 0x52, 0xc4, 0x1c, 0x9e, 0x65, 0x2b, 0x12, 0x8d, 0x18, 0x1a, 0x25, 0x65, 0xff,
	0x97, 0x42, 0x4e, 0x8d, 0x26, 0x3a, 0xd3, 0x58, 0xe4, 0x4d, 0x6e, 0x5f, 0xf6, 0x16, 0xd8, 0xed,
	0x2d, 0xf0, 0xbe, 0xb7, 0xc0, 0xd3, 0xc1, 0xaa, 0xec, 0x0e, 0x56, 0xe5, 0xf5, 0x60, 0x55, 0x1e,
	0x2e, 0x69, 0xa8, 0x96, 0x2b, 0xcf, 0xf1, 0x39, 0x43, 0xe3, 0x28, 0x9a, 0xc5, 0x13, 0xa2, 0xe4,
	0x94, 0x33, 0x84, 0x43, 0x6f, 0x28, 0x78, 0x14, 0xe1, 0x24, 0x41, 0x1b, 0xfd, 0x02, 0xd3, 0x9b,
	0x95, 0x5e, 0x3d, 0x7b, 0x79, 0xd7, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xe9, 0x4a, 0xf6,
	0x11, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	RegisterContract(ctx context.Context, in *MsgRegisterContract, opts ...grpc.CallOption) (*MsgRegisterContractResponse, error)
	DeRegisterContract(ctx context.Context, in *MsgDeRegisterContract, opts ...grpc.CallOption) (*MsgDeRegisterContractResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterContract(ctx context.Context, in *MsgRegisterContract, opts ...grpc.CallOption) (*MsgRegisterContractResponse, error) {
	out := new(MsgRegisterContractResponse)
	err := c.cc.Invoke(ctx, "/aib.cron.v1beta1.Msg/RegisterContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeRegisterContract(ctx context.Context, in *MsgDeRegisterContract, opts ...grpc.CallOption) (*MsgDeRegisterContractResponse, error) {
	out := new(MsgDeRegisterContractResponse)
	err := c.cc.Invoke(ctx, "/aib.cron.v1beta1.Msg/DeRegisterContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	RegisterContract(context.Context, *MsgRegisterContract) (*MsgRegisterContractResponse, error)
	DeRegisterContract(context.Context, *MsgDeRegisterContract) (*MsgDeRegisterContractResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RegisterContract(ctx context.Context, req *MsgRegisterContract) (*MsgRegisterContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterContract not implemented")
}
func (*UnimplementedMsgServer) DeRegisterContract(ctx context.Context, req *MsgDeRegisterContract) (*MsgDeRegisterContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeRegisterContract not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RegisterContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aib.cron.v1beta1.Msg/RegisterContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterContract(ctx, req.(*MsgRegisterContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeRegisterContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeRegisterContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeRegisterContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aib.cron.v1beta1.Msg/DeRegisterContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeRegisterContract(ctx, req.(*MsgDeRegisterContract))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "aib.cron.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterContract",
			Handler:    _Msg_RegisterContract_Handler,
		},
		{
			MethodName: "DeRegisterContract",
			Handler:    _Msg_DeRegisterContract_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aib/cron/v1beta1/tx.proto",
}

func (m *MsgRegisterContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GameType != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GameType))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.GameName) > 0 {
		i -= len(m.GameName)
		copy(dAtA[i:], m.GameName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.GameName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SecurityAddress) > 0 {
		i -= len(m.SecurityAddress)
		copy(dAtA[i:], m.SecurityAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SecurityAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterContractResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterContractResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterContractResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeRegisterContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeRegisterContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeRegisterContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GameId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GameId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.SecurityAddress) > 0 {
		i -= len(m.SecurityAddress)
		copy(dAtA[i:], m.SecurityAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SecurityAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeRegisterContractResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeRegisterContractResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeRegisterContractResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRegisterContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SecurityAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.GameName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.GameType != 0 {
		n += 1 + sovTx(uint64(m.GameType))
	}
	return n
}

func (m *MsgRegisterContractResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeRegisterContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SecurityAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.GameId != 0 {
		n += 1 + sovTx(uint64(m.GameId))
	}
	return n
}

func (m *MsgDeRegisterContractResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRegisterContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegisterContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecurityAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecurityAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GameName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameType", wireType)
			}
			m.GameType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GameType |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRegisterContractResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRegisterContractResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterContractResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDeRegisterContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDeRegisterContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeRegisterContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecurityAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecurityAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameId", wireType)
			}
			m.GameId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GameId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDeRegisterContractResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDeRegisterContractResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeRegisterContractResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
