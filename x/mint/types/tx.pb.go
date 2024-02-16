// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: publicawesome/stargaze/mint/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	math "math"
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

func init() {
	proto.RegisterFile("publicawesome/stargaze/mint/v1beta1/tx.proto", fileDescriptor_fd5f47f6a4b3dad2)
}

var fileDescriptor_fd5f47f6a4b3dad2 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x29, 0x28, 0x4d, 0xca,
	0xc9, 0x4c, 0x4e, 0x2c, 0x4f, 0x2d, 0xce, 0xcf, 0x4d, 0xd5, 0x2f, 0x2e, 0x49, 0x2c, 0x4a, 0x4f,
	0xac, 0x4a, 0xd5, 0xcf, 0xcd, 0xcc, 0x2b, 0xd1, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4,
	0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x46, 0x51, 0xad, 0x07, 0x53, 0xad,
	0x07, 0x52, 0xad, 0x07, 0x55, 0x2d, 0x25, 0x9e, 0x9c, 0x5f, 0x9c, 0x9b, 0x5f, 0xac, 0x9f, 0x5b,
	0x9c, 0xae, 0x5f, 0x66, 0x08, 0xa2, 0x20, 0xba, 0x8d, 0x78, 0xb8, 0x98, 0x7d, 0x8b, 0xd3, 0xa5,
	0x58, 0x1b, 0x9e, 0x6f, 0xd0, 0x62, 0x74, 0xf2, 0x3d, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39,
	0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63,
	0x39, 0x86, 0x28, 0xe3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x88,
	0x85, 0xba, 0x18, 0xee, 0x2b, 0x33, 0x34, 0xd6, 0xaf, 0x80, 0xb8, 0xb2, 0xa4, 0xb2, 0x20, 0xb5,
	0x38, 0x89, 0x0d, 0x6c, 0x87, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x04, 0xe3, 0xdd, 0xfe, 0xd1,
	0x00, 0x00, 0x00,
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
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "publicawesome.stargaze.mint.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "publicawesome/stargaze/mint/v1beta1/tx.proto",
}