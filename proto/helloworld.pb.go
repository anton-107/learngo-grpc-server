// Code generated by protoc-gen-go.
// source: helloworld.proto
// DO NOT EDIT!

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld.proto

It has these top-level messages:
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

type MyMessage struct {
	RequestId string `protobuf:"bytes,1,opt,name=requestId" json:"requestId,omitempty"`
	Text      string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
}

func (m *MyMessage) Reset()                    { *m = MyMessage{} }
func (m *MyMessage) String() string            { return proto.CompactTextString(m) }
func (*MyMessage) ProtoMessage()               {}
func (*MyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*MyMessage)(nil), "MyMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Greeter service

type GreeterClient interface {
	SayWord(ctx context.Context, in *MyMessage, opts ...grpc.CallOption) (*MyMessage, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
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
	SayWord(context.Context, *MyMessage) (*MyMessage, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
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
			MethodName: "SayWord",
			Handler:    _Greeter_SayWord_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x49, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c,
	0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xc8, 0x2a, 0xe9, 0x70, 0x71, 0xfa, 0x56, 0xfa, 0xa6, 0x16, 0x17,
	0x27, 0xa6, 0xa7, 0x0a, 0x09, 0x72, 0x71, 0x16, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x78, 0xa6,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x0a, 0xf1, 0x70, 0xb1, 0x94, 0xa4, 0x56, 0x94, 0x48, 0x30,
	0x81, 0x78, 0x46, 0x1e, 0x5c, 0xec, 0xee, 0x45, 0xa9, 0xa9, 0x25, 0xa9, 0x45, 0x42, 0xb6, 0x5c,
	0xec, 0xc1, 0x89, 0x95, 0xe1, 0xf9, 0x45, 0x29, 0x42, 0x5c, 0x7a, 0x70, 0x23, 0xa4, 0x90, 0xd8,
	0x4a, 0x52, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0x12, 0x11, 0x12, 0xd2, 0x2f, 0x33, 0xd4, 0x47, 0xb8,
	0x4b, 0xbf, 0x38, 0xb1, 0x32, 0x89, 0x0d, 0x6c, 0xbd, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x51,
	0x05, 0x8d, 0x8c, 0xb0, 0x00, 0x00, 0x00,
}
