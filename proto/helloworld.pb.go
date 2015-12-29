// Code generated by protoc-gen-go.
// source: helloworld.proto
// DO NOT EDIT!

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld.proto

It has these top-level messages:
	HealthRequest
	HealthResponse
	MyMessage
*/
package helloworld

import proto "github.com/golang/protobuf/proto"
import "fmt"
import "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HealthRequest struct {
}

func (m *HealthRequest) Reset()                    { *m = HealthRequest{} }
func (m *HealthRequest) String() string            { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()               {}
func (*HealthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HealthResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *HealthResponse) Reset()                    { *m = HealthResponse{} }
func (m *HealthResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()               {}
func (*HealthResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type MyMessage struct {
	RequestId string `protobuf:"bytes,1,opt,name=requestId" json:"requestId,omitempty"`
	Text      string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
}

func (m *MyMessage) Reset()                    { *m = MyMessage{} }
func (m *MyMessage) String() string            { return proto.CompactTextString(m) }
func (*MyMessage) ProtoMessage()               {}
func (*MyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*HealthRequest)(nil), "HealthRequest")
	proto.RegisterType((*HealthResponse)(nil), "HealthResponse")
	proto.RegisterType((*MyMessage)(nil), "MyMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Greeter service

type GreeterClient interface {
	Healthcheck(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	SayWord(ctx context.Context, in *MyMessage, opts ...grpc.CallOption) (*MyMessage, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Healthcheck(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := grpc.Invoke(ctx, "/Greeter/Healthcheck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayWord(ctx context.Context, in *MyMessage, opts ...grpc.CallOption) (*MyMessage, error) {
	out := new(MyMessage)
	err := grpc.Invoke(ctx, "/Greeter/SayWord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	Healthcheck(context.Context, *HealthRequest) (*HealthResponse, error)
	SayWord(context.Context, *MyMessage) (*MyMessage, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Healthcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GreeterServer).Healthcheck(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Greeter_SayWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(MyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GreeterServer).SayWord(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healthcheck",
			Handler:    _Greeter_Healthcheck_Handler,
		},
		{
			MethodName: "SayWord",
			Handler:    _Greeter_SayWord_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x49, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c,
	0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xc8, 0x2a, 0xf1, 0x73, 0xf1, 0x7a, 0xa4, 0x26, 0xe6, 0x94, 0x64,
	0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x28, 0x29, 0x70, 0xf1, 0xc1, 0x04, 0x8a, 0x0b, 0x80,
	0xea, 0x52, 0x85, 0xf8, 0xb8, 0xd8, 0x8a, 0x81, 0x9a, 0x4a, 0x8b, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x95, 0x74, 0xb8, 0x38, 0x7d, 0x2b, 0x7d, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0x04,
	0xb9, 0x38, 0x8b, 0x20, 0x3a, 0x3d, 0x53, 0x20, 0xf2, 0x42, 0x3c, 0x5c, 0x2c, 0x25, 0xa9, 0x15,
	0x25, 0x12, 0x4c, 0x20, 0x9e, 0x51, 0x1f, 0x23, 0x17, 0xbb, 0x7b, 0x51, 0x6a, 0x6a, 0x49, 0x6a,
	0x91, 0x90, 0x0b, 0x17, 0x37, 0xc4, 0xec, 0xe4, 0x8c, 0xd4, 0xe4, 0x6c, 0x21, 0x3e, 0x3d, 0x14,
	0xab, 0xa5, 0xf8, 0xf5, 0x50, 0x6d, 0x56, 0x12, 0x69, 0xba, 0xfc, 0x64, 0x32, 0x13, 0x9f, 0x10,
	0x8f, 0x7e, 0x06, 0x92, 0x36, 0x5b, 0x2e, 0xf6, 0xe0, 0xc4, 0xca, 0xf0, 0xfc, 0xa2, 0x14, 0x21,
	0x2e, 0x3d, 0xb8, 0x4b, 0xa4, 0x90, 0xd8, 0x4a, 0x52, 0x60, 0x8d, 0x22, 0x42, 0x42, 0xfa, 0x65,
	0x86, 0xfa, 0x88, 0x10, 0xd1, 0x2f, 0x4e, 0xac, 0x4c, 0x62, 0x03, 0x7b, 0xdc, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x6c, 0x69, 0x62, 0xcd, 0x2a, 0x01, 0x00, 0x00,
}
