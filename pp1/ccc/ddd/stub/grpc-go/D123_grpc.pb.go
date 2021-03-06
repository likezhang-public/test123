// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package stubs

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// D123Client is the client API for D123 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type D123Client interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type d123Client struct {
	cc grpc.ClientConnInterface
}

func NewD123Client(cc grpc.ClientConnInterface) D123Client {
	return &d123Client{cc}
}

func (c *d123Client) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/ccc.ddd.D123/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// D123Server is the server API for D123 service.
// All implementations must embed UnimplementedD123Server
// for forward compatibility
type D123Server interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedD123Server()
}

// UnimplementedD123Server must be embedded to have forward compatible implementations.
type UnimplementedD123Server struct {
}

func (UnimplementedD123Server) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedD123Server) mustEmbedUnimplementedD123Server() {}

// UnsafeD123Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to D123Server will
// result in compilation errors.
type UnsafeD123Server interface {
	mustEmbedUnimplementedD123Server()
}

func RegisterD123Server(s grpc.ServiceRegistrar, srv D123Server) {
	s.RegisterService(&D123_ServiceDesc, srv)
}

func _D123_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(D123Server).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ccc.ddd.D123/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(D123Server).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// D123_ServiceDesc is the grpc.ServiceDesc for D123 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var D123_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ccc.ddd.D123",
	HandlerType: (*D123Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _D123_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "D123.proto",
}
