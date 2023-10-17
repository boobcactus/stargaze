// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: publicawesome/stargaze/cron/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type MsgPromoteToPrivilegedContract struct {
	// Authority is the address of the governance account or any whitelisted address
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// Contract is the bech32 address of the smart contract
	Contract string `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (m *MsgPromoteToPrivilegedContract) Reset()         { *m = MsgPromoteToPrivilegedContract{} }
func (m *MsgPromoteToPrivilegedContract) String() string { return proto.CompactTextString(m) }
func (*MsgPromoteToPrivilegedContract) ProtoMessage()    {}
func (*MsgPromoteToPrivilegedContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_35ccf9aea279ac0f, []int{0}
}
func (m *MsgPromoteToPrivilegedContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPromoteToPrivilegedContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPromoteToPrivilegedContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPromoteToPrivilegedContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPromoteToPrivilegedContract.Merge(m, src)
}
func (m *MsgPromoteToPrivilegedContract) XXX_Size() int {
	return m.Size()
}
func (m *MsgPromoteToPrivilegedContract) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPromoteToPrivilegedContract.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPromoteToPrivilegedContract proto.InternalMessageInfo

func (m *MsgPromoteToPrivilegedContract) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgPromoteToPrivilegedContract) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

type MsgPromoteToPrivilegedContractResponse struct {
}

func (m *MsgPromoteToPrivilegedContractResponse) Reset() {
	*m = MsgPromoteToPrivilegedContractResponse{}
}
func (m *MsgPromoteToPrivilegedContractResponse) String() string { return proto.CompactTextString(m) }
func (*MsgPromoteToPrivilegedContractResponse) ProtoMessage()    {}
func (*MsgPromoteToPrivilegedContractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_35ccf9aea279ac0f, []int{1}
}
func (m *MsgPromoteToPrivilegedContractResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPromoteToPrivilegedContractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPromoteToPrivilegedContractResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPromoteToPrivilegedContractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPromoteToPrivilegedContractResponse.Merge(m, src)
}
func (m *MsgPromoteToPrivilegedContractResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgPromoteToPrivilegedContractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPromoteToPrivilegedContractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPromoteToPrivilegedContractResponse proto.InternalMessageInfo

type MsgDemoteFromPrivilegedContract struct {
	// Authority is the address of the governance account or any whitelisted address
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// Contract is the bech32 address of the smart contract
	Contract string `protobuf:"bytes,2,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (m *MsgDemoteFromPrivilegedContract) Reset()         { *m = MsgDemoteFromPrivilegedContract{} }
func (m *MsgDemoteFromPrivilegedContract) String() string { return proto.CompactTextString(m) }
func (*MsgDemoteFromPrivilegedContract) ProtoMessage()    {}
func (*MsgDemoteFromPrivilegedContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_35ccf9aea279ac0f, []int{2}
}
func (m *MsgDemoteFromPrivilegedContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDemoteFromPrivilegedContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDemoteFromPrivilegedContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDemoteFromPrivilegedContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDemoteFromPrivilegedContract.Merge(m, src)
}
func (m *MsgDemoteFromPrivilegedContract) XXX_Size() int {
	return m.Size()
}
func (m *MsgDemoteFromPrivilegedContract) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDemoteFromPrivilegedContract.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDemoteFromPrivilegedContract proto.InternalMessageInfo

func (m *MsgDemoteFromPrivilegedContract) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgDemoteFromPrivilegedContract) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

type MsgDemoteFromPrivilegedContractResponse struct {
}

func (m *MsgDemoteFromPrivilegedContractResponse) Reset() {
	*m = MsgDemoteFromPrivilegedContractResponse{}
}
func (m *MsgDemoteFromPrivilegedContractResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDemoteFromPrivilegedContractResponse) ProtoMessage()    {}
func (*MsgDemoteFromPrivilegedContractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_35ccf9aea279ac0f, []int{3}
}
func (m *MsgDemoteFromPrivilegedContractResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDemoteFromPrivilegedContractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDemoteFromPrivilegedContractResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDemoteFromPrivilegedContractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDemoteFromPrivilegedContractResponse.Merge(m, src)
}
func (m *MsgDemoteFromPrivilegedContractResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDemoteFromPrivilegedContractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDemoteFromPrivilegedContractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDemoteFromPrivilegedContractResponse proto.InternalMessageInfo

type MsgUpdateParams struct {
	// Authority is the address of the governance account.
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_35ccf9aea279ac0f, []int{4}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_35ccf9aea279ac0f, []int{5}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgPromoteToPrivilegedContract)(nil), "publicawesome.stargaze.cron.v1.MsgPromoteToPrivilegedContract")
	proto.RegisterType((*MsgPromoteToPrivilegedContractResponse)(nil), "publicawesome.stargaze.cron.v1.MsgPromoteToPrivilegedContractResponse")
	proto.RegisterType((*MsgDemoteFromPrivilegedContract)(nil), "publicawesome.stargaze.cron.v1.MsgDemoteFromPrivilegedContract")
	proto.RegisterType((*MsgDemoteFromPrivilegedContractResponse)(nil), "publicawesome.stargaze.cron.v1.MsgDemoteFromPrivilegedContractResponse")
	proto.RegisterType((*MsgUpdateParams)(nil), "publicawesome.stargaze.cron.v1.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "publicawesome.stargaze.cron.v1.MsgUpdateParamsResponse")
}

func init() {
	proto.RegisterFile("publicawesome/stargaze/cron/v1/tx.proto", fileDescriptor_35ccf9aea279ac0f)
}

var fileDescriptor_35ccf9aea279ac0f = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x4d, 0x8b, 0xd3, 0x40,
	0x18, 0xc7, 0x33, 0x2e, 0x2c, 0xee, 0x28, 0x2a, 0x41, 0xd8, 0x1a, 0x97, 0xa9, 0xf4, 0xb0, 0x2f,
	0xa2, 0x19, 0xda, 0x1e, 0x04, 0x0f, 0x0a, 0x55, 0xea, 0x29, 0x50, 0x8a, 0x82, 0x78, 0x9b, 0xa6,
	0xc3, 0x34, 0xda, 0xe9, 0x13, 0x66, 0xa6, 0xb1, 0xf5, 0xe8, 0x27, 0xf0, 0xea, 0x07, 0x10, 0x3c,
	0xfa, 0x31, 0x7a, 0x2c, 0x9e, 0x3c, 0x15, 0x69, 0x0f, 0x82, 0x9f, 0x42, 0x92, 0x34, 0xad, 0xb5,
	0x90, 0x54, 0xca, 0xde, 0x92, 0xcc, 0xff, 0xe5, 0x17, 0x9e, 0xe1, 0xc1, 0x67, 0xe1, 0xb0, 0xd3,
	0x0f, 0x7c, 0xf6, 0x9e, 0x6b, 0x90, 0x9c, 0x6a, 0xc3, 0x94, 0x60, 0x1f, 0x38, 0xf5, 0x15, 0x0c,
	0x68, 0x54, 0xa5, 0x66, 0xe4, 0x86, 0x0a, 0x0c, 0xd8, 0x64, 0x43, 0xe8, 0x66, 0x42, 0x37, 0x16,
	0xba, 0x51, 0xd5, 0xb9, 0x2d, 0x40, 0x40, 0x22, 0xa5, 0xf1, 0x53, 0xea, 0x72, 0x8e, 0x7d, 0xd0,
	0x12, 0x34, 0x95, 0x5a, 0xc4, 0x69, 0x52, 0x8b, 0xe5, 0xc1, 0x45, 0x41, 0x6f, 0x12, 0x9b, 0x48,
	0x2b, 0x6f, 0x31, 0xf1, 0xb4, 0x68, 0x29, 0x90, 0x60, 0xf8, 0x4b, 0x68, 0xa9, 0x20, 0x0a, 0xfa,
	0x5c, 0xf0, 0xee, 0x33, 0x18, 0x18, 0xc5, 0x7c, 0x63, 0x9f, 0xe0, 0x23, 0x36, 0x34, 0x3d, 0x50,
	0x81, 0x19, 0x97, 0xd0, 0x3d, 0x74, 0x7e, 0xd4, 0x5e, 0x7f, 0xb0, 0x1d, 0x7c, 0xd5, 0x5f, 0x2a,
	0x4b, 0x57, 0x92, 0xc3, 0xd5, 0xfb, 0xe3, 0x1b, 0x1f, 0x7f, 0x7d, 0xbb, 0xbf, 0xd6, 0x56, 0xce,
	0xf1, 0x69, 0x7e, 0x57, 0x9b, 0xeb, 0x10, 0x06, 0x9a, 0x57, 0xde, 0xe1, 0xb2, 0xa7, 0xc5, 0x73,
	0x1e, 0x0b, 0x9b, 0x0a, 0xe4, 0xa5, 0x62, 0x5d, 0xe0, 0xb3, 0x82, 0xb2, 0x15, 0xd7, 0x67, 0x84,
	0x6f, 0x7a, 0x5a, 0xbc, 0x0a, 0xbb, 0xcc, 0xf0, 0x16, 0x53, 0x4c, 0xea, 0x02, 0x90, 0xd7, 0xf8,
	0x30, 0x4c, 0x74, 0x09, 0xc6, 0xb5, 0xda, 0xa9, 0x9b, 0x3f, 0x6a, 0x37, 0x4d, 0x6d, 0x94, 0x26,
	0xb3, 0xb2, 0xf5, 0x7b, 0x56, 0xbe, 0x95, 0xba, 0x1f, 0x80, 0x0c, 0x0c, 0x97, 0xa1, 0x19, 0xb7,
	0x97, 0x79, 0x5b, 0xbf, 0x71, 0x07, 0x1f, 0xff, 0x83, 0x96, 0x61, 0xd7, 0xbe, 0x1f, 0xe0, 0x03,
	0x4f, 0x0b, 0xfb, 0x0b, 0xc2, 0x77, 0xf3, 0x46, 0xfd, 0xa4, 0x08, 0x2e, 0x7f, 0x7c, 0x4e, 0x73,
	0x3f, 0x7f, 0xc6, 0x6b, 0x7f, 0x45, 0xf8, 0x24, 0x77, 0xf8, 0x4f, 0x77, 0x28, 0xca, 0x0b, 0x70,
	0x5e, 0xec, 0x19, 0xb0, 0x42, 0x1d, 0xe1, 0xeb, 0x1b, 0xb7, 0x81, 0xee, 0x10, 0xfc, 0xb7, 0xc1,
	0x79, 0xf4, 0x9f, 0x86, 0xac, 0xb9, 0xe1, 0x4d, 0xe6, 0x04, 0x4d, 0xe7, 0x04, 0xfd, 0x9c, 0x13,
	0xf4, 0x69, 0x41, 0xac, 0xe9, 0x82, 0x58, 0x3f, 0x16, 0xc4, 0x7a, 0x53, 0x17, 0x81, 0xe9, 0x0d,
	0x3b, 0xae, 0x0f, 0x92, 0xa6, 0xe1, 0x0f, 0xb7, 0x56, 0x41, 0x54, 0xad, 0xd1, 0x51, 0xba, 0x10,
	0xcc, 0x38, 0xe4, 0xba, 0x73, 0x98, 0xec, 0x83, 0xfa, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9,
	0x3c, 0x43, 0xac, 0xb4, 0x04, 0x00, 0x00,
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
	PromoteToPrivilegedContract(ctx context.Context, in *MsgPromoteToPrivilegedContract, opts ...grpc.CallOption) (*MsgPromoteToPrivilegedContractResponse, error)
	DemoteFromPrivilegedContract(ctx context.Context, in *MsgDemoteFromPrivilegedContract, opts ...grpc.CallOption) (*MsgDemoteFromPrivilegedContractResponse, error)
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) PromoteToPrivilegedContract(ctx context.Context, in *MsgPromoteToPrivilegedContract, opts ...grpc.CallOption) (*MsgPromoteToPrivilegedContractResponse, error) {
	out := new(MsgPromoteToPrivilegedContractResponse)
	err := c.cc.Invoke(ctx, "/publicawesome.stargaze.cron.v1.Msg/PromoteToPrivilegedContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DemoteFromPrivilegedContract(ctx context.Context, in *MsgDemoteFromPrivilegedContract, opts ...grpc.CallOption) (*MsgDemoteFromPrivilegedContractResponse, error) {
	out := new(MsgDemoteFromPrivilegedContractResponse)
	err := c.cc.Invoke(ctx, "/publicawesome.stargaze.cron.v1.Msg/DemoteFromPrivilegedContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/publicawesome.stargaze.cron.v1.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	PromoteToPrivilegedContract(context.Context, *MsgPromoteToPrivilegedContract) (*MsgPromoteToPrivilegedContractResponse, error)
	DemoteFromPrivilegedContract(context.Context, *MsgDemoteFromPrivilegedContract) (*MsgDemoteFromPrivilegedContractResponse, error)
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) PromoteToPrivilegedContract(ctx context.Context, req *MsgPromoteToPrivilegedContract) (*MsgPromoteToPrivilegedContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromoteToPrivilegedContract not implemented")
}
func (*UnimplementedMsgServer) DemoteFromPrivilegedContract(ctx context.Context, req *MsgDemoteFromPrivilegedContract) (*MsgDemoteFromPrivilegedContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DemoteFromPrivilegedContract not implemented")
}
func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_PromoteToPrivilegedContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPromoteToPrivilegedContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PromoteToPrivilegedContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicawesome.stargaze.cron.v1.Msg/PromoteToPrivilegedContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PromoteToPrivilegedContract(ctx, req.(*MsgPromoteToPrivilegedContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DemoteFromPrivilegedContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDemoteFromPrivilegedContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DemoteFromPrivilegedContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicawesome.stargaze.cron.v1.Msg/DemoteFromPrivilegedContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DemoteFromPrivilegedContract(ctx, req.(*MsgDemoteFromPrivilegedContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicawesome.stargaze.cron.v1.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "publicawesome.stargaze.cron.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PromoteToPrivilegedContract",
			Handler:    _Msg_PromoteToPrivilegedContract_Handler,
		},
		{
			MethodName: "DemoteFromPrivilegedContract",
			Handler:    _Msg_DemoteFromPrivilegedContract_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "publicawesome/stargaze/cron/v1/tx.proto",
}

func (m *MsgPromoteToPrivilegedContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPromoteToPrivilegedContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPromoteToPrivilegedContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgPromoteToPrivilegedContractResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPromoteToPrivilegedContractResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPromoteToPrivilegedContractResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDemoteFromPrivilegedContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDemoteFromPrivilegedContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDemoteFromPrivilegedContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDemoteFromPrivilegedContractResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDemoteFromPrivilegedContractResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDemoteFromPrivilegedContractResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgPromoteToPrivilegedContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgPromoteToPrivilegedContractResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDemoteFromPrivilegedContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDemoteFromPrivilegedContractResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
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
func (m *MsgPromoteToPrivilegedContract) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgPromoteToPrivilegedContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPromoteToPrivilegedContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
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
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
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
			m.Contract = string(dAtA[iNdEx:postIndex])
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
func (m *MsgPromoteToPrivilegedContractResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgPromoteToPrivilegedContractResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPromoteToPrivilegedContractResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgDemoteFromPrivilegedContract) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDemoteFromPrivilegedContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDemoteFromPrivilegedContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
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
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
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
			m.Contract = string(dAtA[iNdEx:postIndex])
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
func (m *MsgDemoteFromPrivilegedContractResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDemoteFromPrivilegedContractResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDemoteFromPrivilegedContractResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
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
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
