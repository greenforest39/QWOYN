// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stimulus/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type MsgDepositIntoOutpostFunding struct {
	Sender string     `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Amount types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgDepositIntoOutpostFunding) Reset()         { *m = MsgDepositIntoOutpostFunding{} }
func (m *MsgDepositIntoOutpostFunding) String() string { return proto.CompactTextString(m) }
func (*MsgDepositIntoOutpostFunding) ProtoMessage()    {}
func (*MsgDepositIntoOutpostFunding) Descriptor() ([]byte, []int) {
	return fileDescriptor_998c0a138c8bb0b7, []int{0}
}
func (m *MsgDepositIntoOutpostFunding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDepositIntoOutpostFunding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDepositIntoOutpostFunding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDepositIntoOutpostFunding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDepositIntoOutpostFunding.Merge(m, src)
}
func (m *MsgDepositIntoOutpostFunding) XXX_Size() int {
	return m.Size()
}
func (m *MsgDepositIntoOutpostFunding) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDepositIntoOutpostFunding.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDepositIntoOutpostFunding proto.InternalMessageInfo

func (m *MsgDepositIntoOutpostFunding) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgDepositIntoOutpostFunding) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

type MsgDepositIntoOutpostFundingResponse struct {
}

func (m *MsgDepositIntoOutpostFundingResponse) Reset()         { *m = MsgDepositIntoOutpostFundingResponse{} }
func (m *MsgDepositIntoOutpostFundingResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDepositIntoOutpostFundingResponse) ProtoMessage()    {}
func (*MsgDepositIntoOutpostFundingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_998c0a138c8bb0b7, []int{1}
}
func (m *MsgDepositIntoOutpostFundingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDepositIntoOutpostFundingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDepositIntoOutpostFundingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDepositIntoOutpostFundingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDepositIntoOutpostFundingResponse.Merge(m, src)
}
func (m *MsgDepositIntoOutpostFundingResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDepositIntoOutpostFundingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDepositIntoOutpostFundingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDepositIntoOutpostFundingResponse proto.InternalMessageInfo

type MsgWithdrawFromOutpostFunding struct {
	Sender string     `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Amount types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgWithdrawFromOutpostFunding) Reset()         { *m = MsgWithdrawFromOutpostFunding{} }
func (m *MsgWithdrawFromOutpostFunding) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawFromOutpostFunding) ProtoMessage()    {}
func (*MsgWithdrawFromOutpostFunding) Descriptor() ([]byte, []int) {
	return fileDescriptor_998c0a138c8bb0b7, []int{2}
}
func (m *MsgWithdrawFromOutpostFunding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawFromOutpostFunding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawFromOutpostFunding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawFromOutpostFunding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawFromOutpostFunding.Merge(m, src)
}
func (m *MsgWithdrawFromOutpostFunding) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawFromOutpostFunding) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawFromOutpostFunding.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawFromOutpostFunding proto.InternalMessageInfo

func (m *MsgWithdrawFromOutpostFunding) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgWithdrawFromOutpostFunding) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

type MsgWithdrawFromOutpostFundingResponse struct {
}

func (m *MsgWithdrawFromOutpostFundingResponse) Reset()         { *m = MsgWithdrawFromOutpostFundingResponse{} }
func (m *MsgWithdrawFromOutpostFundingResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawFromOutpostFundingResponse) ProtoMessage()    {}
func (*MsgWithdrawFromOutpostFundingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_998c0a138c8bb0b7, []int{3}
}
func (m *MsgWithdrawFromOutpostFundingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawFromOutpostFundingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawFromOutpostFundingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawFromOutpostFundingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawFromOutpostFundingResponse.Merge(m, src)
}
func (m *MsgWithdrawFromOutpostFundingResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawFromOutpostFundingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawFromOutpostFundingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawFromOutpostFundingResponse proto.InternalMessageInfo

type EventDepositIntoOutpostFunding struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *EventDepositIntoOutpostFunding) Reset()         { *m = EventDepositIntoOutpostFunding{} }
func (m *EventDepositIntoOutpostFunding) String() string { return proto.CompactTextString(m) }
func (*EventDepositIntoOutpostFunding) ProtoMessage()    {}
func (*EventDepositIntoOutpostFunding) Descriptor() ([]byte, []int) {
	return fileDescriptor_998c0a138c8bb0b7, []int{4}
}
func (m *EventDepositIntoOutpostFunding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventDepositIntoOutpostFunding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventDepositIntoOutpostFunding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventDepositIntoOutpostFunding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventDepositIntoOutpostFunding.Merge(m, src)
}
func (m *EventDepositIntoOutpostFunding) XXX_Size() int {
	return m.Size()
}
func (m *EventDepositIntoOutpostFunding) XXX_DiscardUnknown() {
	xxx_messageInfo_EventDepositIntoOutpostFunding.DiscardUnknown(m)
}

var xxx_messageInfo_EventDepositIntoOutpostFunding proto.InternalMessageInfo

func (m *EventDepositIntoOutpostFunding) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *EventDepositIntoOutpostFunding) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

type EventWithdrawFromOutpostFunding struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *EventWithdrawFromOutpostFunding) Reset()         { *m = EventWithdrawFromOutpostFunding{} }
func (m *EventWithdrawFromOutpostFunding) String() string { return proto.CompactTextString(m) }
func (*EventWithdrawFromOutpostFunding) ProtoMessage()    {}
func (*EventWithdrawFromOutpostFunding) Descriptor() ([]byte, []int) {
	return fileDescriptor_998c0a138c8bb0b7, []int{5}
}
func (m *EventWithdrawFromOutpostFunding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventWithdrawFromOutpostFunding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventWithdrawFromOutpostFunding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventWithdrawFromOutpostFunding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventWithdrawFromOutpostFunding.Merge(m, src)
}
func (m *EventWithdrawFromOutpostFunding) XXX_Size() int {
	return m.Size()
}
func (m *EventWithdrawFromOutpostFunding) XXX_DiscardUnknown() {
	xxx_messageInfo_EventWithdrawFromOutpostFunding.DiscardUnknown(m)
}

var xxx_messageInfo_EventWithdrawFromOutpostFunding proto.InternalMessageInfo

func (m *EventWithdrawFromOutpostFunding) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *EventWithdrawFromOutpostFunding) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgDepositIntoOutpostFunding)(nil), "cosmichorizon.qwoyn.stimulus.MsgDepositIntoOutpostFunding")
	proto.RegisterType((*MsgDepositIntoOutpostFundingResponse)(nil), "cosmichorizon.qwoyn.stimulus.MsgDepositIntoOutpostFundingResponse")
	proto.RegisterType((*MsgWithdrawFromOutpostFunding)(nil), "cosmichorizon.qwoyn.stimulus.MsgWithdrawFromOutpostFunding")
	proto.RegisterType((*MsgWithdrawFromOutpostFundingResponse)(nil), "cosmichorizon.qwoyn.stimulus.MsgWithdrawFromOutpostFundingResponse")
	proto.RegisterType((*EventDepositIntoOutpostFunding)(nil), "cosmichorizon.qwoyn.stimulus.EventDepositIntoOutpostFunding")
	proto.RegisterType((*EventWithdrawFromOutpostFunding)(nil), "cosmichorizon.qwoyn.stimulus.EventWithdrawFromOutpostFunding")
}

func init() { proto.RegisterFile("stimulus/tx.proto", fileDescriptor_998c0a138c8bb0b7) }

var fileDescriptor_998c0a138c8bb0b7 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xcf, 0x6b, 0xe2, 0x40,
	0x14, 0xc7, 0x33, 0xee, 0x22, 0x38, 0x7b, 0xda, 0xb0, 0x88, 0x06, 0x77, 0x14, 0xd9, 0x1f, 0x5e,
	0x76, 0x66, 0x75, 0x0f, 0x0b, 0xbb, 0x37, 0x6d, 0x85, 0x16, 0xa4, 0xad, 0x97, 0x42, 0x6f, 0x49,
	0x1c, 0xe2, 0x40, 0x33, 0x2f, 0xcd, 0x4c, 0xfc, 0xd1, 0xbf, 0xa2, 0xf7, 0x42, 0xff, 0x9a, 0x1e,
	0x3c, 0x7a, 0xec, 0xa9, 0x14, 0xfd, 0x47, 0x8a, 0x49, 0xb4, 0xf4, 0x60, 0x4a, 0x2d, 0xbd, 0xbd,
	0x64, 0xde, 0xf7, 0x7d, 0xbe, 0xef, 0xcd, 0x3c, 0xfc, 0x59, 0x69, 0xe1, 0x47, 0xe7, 0x91, 0x62,
	0x7a, 0x42, 0x83, 0x10, 0x34, 0x98, 0x15, 0x17, 0x94, 0x2f, 0xdc, 0x21, 0x84, 0xe2, 0x12, 0x24,
	0xbd, 0x18, 0xc3, 0x54, 0xd2, 0x75, 0x9a, 0x45, 0x56, 0xa7, 0xa0, 0x98, 0x63, 0x2b, 0xce, 0x46,
	0x4d, 0x87, 0x6b, 0xbb, 0xc9, 0x5c, 0x10, 0x32, 0x51, 0x5b, 0x5f, 0x3c, 0xf0, 0x20, 0x0e, 0xd9,
	0x2a, 0x4a, 0xfe, 0xd6, 0x01, 0x57, 0x7a, 0xca, 0xdb, 0xe3, 0x01, 0x28, 0xa1, 0x0f, 0xa4, 0x86,
	0xa3, 0x48, 0x07, 0xa0, 0x74, 0x37, 0x92, 0x03, 0x21, 0x3d, 0xb3, 0x88, 0xf3, 0x8a, 0xcb, 0x01,
	0x0f, 0x4b, 0xa8, 0x86, 0x1a, 0x85, 0x7e, 0xfa, 0x65, 0xfe, 0xc5, 0x79, 0xdb, 0x87, 0x48, 0xea,
	0x52, 0xae, 0x86, 0x1a, 0x9f, 0x5a, 0x65, 0x9a, 0xe0, 0xe9, 0x0a, 0x4f, 0x53, 0x3c, 0xed, 0x80,
	0x90, 0xed, 0x8f, 0xb3, 0xfb, 0xaa, 0xd1, 0x4f, 0xd3, 0xeb, 0x3f, 0xf0, 0xb7, 0x2c, 0x60, 0x9f,
	0xab, 0x00, 0xa4, 0xe2, 0xf5, 0x00, 0x7f, 0xed, 0x29, 0xef, 0x54, 0xe8, 0xe1, 0x20, 0xb4, 0xc7,
	0xdd, 0x10, 0xfc, 0xf7, 0x76, 0xf6, 0x13, 0x7f, 0xcf, 0x24, 0x6e, 0xac, 0x1d, 0x63, 0xb2, 0x3f,
	0xe2, 0x52, 0xbf, 0x7e, 0x6a, 0xc5, 0x67, 0xde, 0x0a, 0x1b, 0xf4, 0x09, 0xae, 0xc6, 0x15, 0x77,
	0x68, 0x77, 0x4b, 0xc9, 0xd6, 0x6d, 0x0e, 0x7f, 0xe8, 0x29, 0xcf, 0xbc, 0x46, 0xb8, 0xbc, 0xdd,
	0xe8, 0x3f, 0x9a, 0xf5, 0xa6, 0x68, 0xd6, 0x4d, 0x59, 0xed, 0xdd, 0xb5, 0xeb, 0x51, 0x9a, 0x37,
	0x08, 0x5b, 0x19, 0x4d, 0xff, 0x7f, 0x11, 0xb1, 0x5d, 0x6c, 0x75, 0xde, 0x20, 0x5e, 0x1b, 0x6c,
	0x1f, 0xce, 0x16, 0x04, 0xcd, 0x17, 0x04, 0x3d, 0x2c, 0x08, 0xba, 0x5a, 0x12, 0x63, 0xbe, 0x24,
	0xc6, 0xdd, 0x92, 0x18, 0x67, 0xbf, 0x3d, 0xa1, 0x87, 0x91, 0x43, 0x5d, 0xf0, 0x59, 0x02, 0xfa,
	0x95, 0x92, 0x58, 0x4c, 0x62, 0x13, 0xf6, 0xb4, 0xc2, 0xd3, 0x80, 0x2b, 0x27, 0x1f, 0xaf, 0xdc,
	0x9f, 0xc7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xd0, 0x31, 0xf8, 0xdb, 0x03, 0x00, 0x00,
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
	DepositIntoOutpostFunding(ctx context.Context, in *MsgDepositIntoOutpostFunding, opts ...grpc.CallOption) (*MsgDepositIntoOutpostFundingResponse, error)
	WithdrawFromOutpostFunding(ctx context.Context, in *MsgWithdrawFromOutpostFunding, opts ...grpc.CallOption) (*MsgWithdrawFromOutpostFundingResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) DepositIntoOutpostFunding(ctx context.Context, in *MsgDepositIntoOutpostFunding, opts ...grpc.CallOption) (*MsgDepositIntoOutpostFundingResponse, error) {
	out := new(MsgDepositIntoOutpostFundingResponse)
	err := c.cc.Invoke(ctx, "/cosmichorizon.qwoyn.stimulus.Msg/DepositIntoOutpostFunding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawFromOutpostFunding(ctx context.Context, in *MsgWithdrawFromOutpostFunding, opts ...grpc.CallOption) (*MsgWithdrawFromOutpostFundingResponse, error) {
	out := new(MsgWithdrawFromOutpostFundingResponse)
	err := c.cc.Invoke(ctx, "/cosmichorizon.qwoyn.stimulus.Msg/WithdrawFromOutpostFunding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	DepositIntoOutpostFunding(context.Context, *MsgDepositIntoOutpostFunding) (*MsgDepositIntoOutpostFundingResponse, error)
	WithdrawFromOutpostFunding(context.Context, *MsgWithdrawFromOutpostFunding) (*MsgWithdrawFromOutpostFundingResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) DepositIntoOutpostFunding(ctx context.Context, req *MsgDepositIntoOutpostFunding) (*MsgDepositIntoOutpostFundingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DepositIntoOutpostFunding not implemented")
}
func (*UnimplementedMsgServer) WithdrawFromOutpostFunding(ctx context.Context, req *MsgWithdrawFromOutpostFunding) (*MsgWithdrawFromOutpostFundingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawFromOutpostFunding not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_DepositIntoOutpostFunding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDepositIntoOutpostFunding)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DepositIntoOutpostFunding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmichorizon.qwoyn.stimulus.Msg/DepositIntoOutpostFunding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DepositIntoOutpostFunding(ctx, req.(*MsgDepositIntoOutpostFunding))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawFromOutpostFunding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawFromOutpostFunding)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawFromOutpostFunding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmichorizon.qwoyn.stimulus.Msg/WithdrawFromOutpostFunding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawFromOutpostFunding(ctx, req.(*MsgWithdrawFromOutpostFunding))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmichorizon.qwoyn.stimulus.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DepositIntoOutpostFunding",
			Handler:    _Msg_DepositIntoOutpostFunding_Handler,
		},
		{
			MethodName: "WithdrawFromOutpostFunding",
			Handler:    _Msg_WithdrawFromOutpostFunding_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stimulus/tx.proto",
}

func (m *MsgDepositIntoOutpostFunding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDepositIntoOutpostFunding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDepositIntoOutpostFunding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDepositIntoOutpostFundingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDepositIntoOutpostFundingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDepositIntoOutpostFundingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawFromOutpostFunding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawFromOutpostFunding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawFromOutpostFunding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawFromOutpostFundingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawFromOutpostFundingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawFromOutpostFundingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *EventDepositIntoOutpostFunding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventDepositIntoOutpostFunding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventDepositIntoOutpostFunding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventWithdrawFromOutpostFunding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventWithdrawFromOutpostFunding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventWithdrawFromOutpostFunding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
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
func (m *MsgDepositIntoOutpostFunding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgDepositIntoOutpostFundingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgWithdrawFromOutpostFunding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgWithdrawFromOutpostFundingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *EventDepositIntoOutpostFunding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *EventWithdrawFromOutpostFunding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgDepositIntoOutpostFunding) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDepositIntoOutpostFunding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDepositIntoOutpostFunding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *MsgDepositIntoOutpostFundingResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDepositIntoOutpostFundingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDepositIntoOutpostFundingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgWithdrawFromOutpostFunding) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgWithdrawFromOutpostFunding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawFromOutpostFunding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *MsgWithdrawFromOutpostFundingResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgWithdrawFromOutpostFundingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawFromOutpostFundingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *EventDepositIntoOutpostFunding) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventDepositIntoOutpostFunding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventDepositIntoOutpostFunding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
func (m *EventWithdrawFromOutpostFunding) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventWithdrawFromOutpostFunding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventWithdrawFromOutpostFunding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
