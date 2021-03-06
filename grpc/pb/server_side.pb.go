// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/server_side.proto

package pb // import "./pb"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ServerSideDeliveryRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerSideDeliveryRequest) Reset()         { *m = ServerSideDeliveryRequest{} }
func (m *ServerSideDeliveryRequest) String() string { return proto.CompactTextString(m) }
func (*ServerSideDeliveryRequest) ProtoMessage()    {}
func (*ServerSideDeliveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_side_7f76af0eaeefdf2f, []int{0}
}
func (m *ServerSideDeliveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerSideDeliveryRequest.Unmarshal(m, b)
}
func (m *ServerSideDeliveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerSideDeliveryRequest.Marshal(b, m, deterministic)
}
func (dst *ServerSideDeliveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerSideDeliveryRequest.Merge(dst, src)
}
func (m *ServerSideDeliveryRequest) XXX_Size() int {
	return xxx_messageInfo_ServerSideDeliveryRequest.Size(m)
}
func (m *ServerSideDeliveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerSideDeliveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerSideDeliveryRequest proto.InternalMessageInfo

func (m *ServerSideDeliveryRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type ServerSideDeliveryResponse struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerSideDeliveryResponse) Reset()         { *m = ServerSideDeliveryResponse{} }
func (m *ServerSideDeliveryResponse) String() string { return proto.CompactTextString(m) }
func (*ServerSideDeliveryResponse) ProtoMessage()    {}
func (*ServerSideDeliveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_side_7f76af0eaeefdf2f, []int{1}
}
func (m *ServerSideDeliveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerSideDeliveryResponse.Unmarshal(m, b)
}
func (m *ServerSideDeliveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerSideDeliveryResponse.Marshal(b, m, deterministic)
}
func (dst *ServerSideDeliveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerSideDeliveryResponse.Merge(dst, src)
}
func (m *ServerSideDeliveryResponse) XXX_Size() int {
	return xxx_messageInfo_ServerSideDeliveryResponse.Size(m)
}
func (m *ServerSideDeliveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerSideDeliveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerSideDeliveryResponse proto.InternalMessageInfo

func (m *ServerSideDeliveryResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*ServerSideDeliveryRequest)(nil), "proto.ServerSideDeliveryRequest")
	proto.RegisterType((*ServerSideDeliveryResponse)(nil), "proto.ServerSideDeliveryResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServerSideDeliveryClient is the client API for ServerSideDelivery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerSideDeliveryClient interface {
	Watch(ctx context.Context, in *ServerSideDeliveryRequest, opts ...grpc.CallOption) (ServerSideDelivery_WatchClient, error)
}

type serverSideDeliveryClient struct {
	cc *grpc.ClientConn
}

func NewServerSideDeliveryClient(cc *grpc.ClientConn) ServerSideDeliveryClient {
	return &serverSideDeliveryClient{cc}
}

func (c *serverSideDeliveryClient) Watch(ctx context.Context, in *ServerSideDeliveryRequest, opts ...grpc.CallOption) (ServerSideDelivery_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ServerSideDelivery_serviceDesc.Streams[0], "/proto.ServerSideDelivery/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &serverSideDeliveryWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServerSideDelivery_WatchClient interface {
	Recv() (*ServerSideDeliveryResponse, error)
	grpc.ClientStream
}

type serverSideDeliveryWatchClient struct {
	grpc.ClientStream
}

func (x *serverSideDeliveryWatchClient) Recv() (*ServerSideDeliveryResponse, error) {
	m := new(ServerSideDeliveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerSideDeliveryServer is the server API for ServerSideDelivery service.
type ServerSideDeliveryServer interface {
	Watch(*ServerSideDeliveryRequest, ServerSideDelivery_WatchServer) error
}

func RegisterServerSideDeliveryServer(s *grpc.Server, srv ServerSideDeliveryServer) {
	s.RegisterService(&_ServerSideDelivery_serviceDesc, srv)
}

func _ServerSideDelivery_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ServerSideDeliveryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServerSideDeliveryServer).Watch(m, &serverSideDeliveryWatchServer{stream})
}

type ServerSideDelivery_WatchServer interface {
	Send(*ServerSideDeliveryResponse) error
	grpc.ServerStream
}

type serverSideDeliveryWatchServer struct {
	grpc.ServerStream
}

func (x *serverSideDeliveryWatchServer) Send(m *ServerSideDeliveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ServerSideDelivery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ServerSideDelivery",
	HandlerType: (*ServerSideDeliveryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _ServerSideDelivery_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/server_side.proto",
}

func init() {
	proto.RegisterFile("proto/server_side.proto", fileDescriptor_server_side_7f76af0eaeefdf2f)
}

var fileDescriptor_server_side_7f76af0eaeefdf2f = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x8a, 0x2f, 0xce, 0x4c, 0x49, 0xd5, 0x03, 0x8b,
	0x08, 0xb1, 0x82, 0x29, 0x25, 0x5d, 0x2e, 0xc9, 0x60, 0xb0, 0x5c, 0x70, 0x66, 0x4a, 0xaa, 0x4b,
	0x6a, 0x4e, 0x66, 0x59, 0x6a, 0x51, 0x65, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x00,
	0x17, 0x73, 0x76, 0x6a, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x88, 0xa9, 0xa4, 0xc7,
	0x25, 0x85, 0x4d, 0x79, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0xa6, 0x7a, 0xa3, 0x34, 0x2e, 0x21,
	0x4c, 0xf5, 0x42, 0x01, 0x5c, 0xac, 0xe1, 0x89, 0x25, 0xc9, 0x19, 0x42, 0x0a, 0x10, 0xc7, 0xe8,
	0xe1, 0x74, 0x82, 0x94, 0x22, 0x1e, 0x15, 0x10, 0x5b, 0x95, 0x18, 0x0c, 0x18, 0x9d, 0x38, 0xa3,
	0xd8, 0xf5, 0xf4, 0x0b, 0x92, 0xac, 0x0b, 0x92, 0x92, 0xd8, 0xc0, 0x1a, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x4d, 0xa2, 0x5a, 0x3e, 0xfa, 0x00, 0x00, 0x00,
}
