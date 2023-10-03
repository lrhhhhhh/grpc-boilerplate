// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: proto/hello.proto

package helloService

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

const (
	Hello_Hello_FullMethodName                   = "/helloService.Hello/Hello"
	Hello_ClientStreamPing_FullMethodName        = "/helloService.Hello/ClientStreamPing"
	Hello_ServerStreamPing_FullMethodName        = "/helloService.Hello/ServerStreamPing"
	Hello_BiDirectionalStreamPing_FullMethodName = "/helloService.Hello/BiDirectionalStreamPing"
)

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	ClientStreamPing(ctx context.Context, opts ...grpc.CallOption) (Hello_ClientStreamPingClient, error)
	ServerStreamPing(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (Hello_ServerStreamPingClient, error)
	BiDirectionalStreamPing(ctx context.Context, opts ...grpc.CallOption) (Hello_BiDirectionalStreamPingClient, error)
}

type helloClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloClient(cc grpc.ClientConnInterface) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, Hello_Hello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloClient) ClientStreamPing(ctx context.Context, opts ...grpc.CallOption) (Hello_ClientStreamPingClient, error) {
	stream, err := c.cc.NewStream(ctx, &Hello_ServiceDesc.Streams[0], Hello_ClientStreamPing_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &helloClientStreamPingClient{stream}
	return x, nil
}

type Hello_ClientStreamPingClient interface {
	Send(*PingRequest) error
	CloseAndRecv() (*PingReply, error)
	grpc.ClientStream
}

type helloClientStreamPingClient struct {
	grpc.ClientStream
}

func (x *helloClientStreamPingClient) Send(m *PingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloClientStreamPingClient) CloseAndRecv() (*PingReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PingReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloClient) ServerStreamPing(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (Hello_ServerStreamPingClient, error) {
	stream, err := c.cc.NewStream(ctx, &Hello_ServiceDesc.Streams[1], Hello_ServerStreamPing_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServerStreamPingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Hello_ServerStreamPingClient interface {
	Recv() (*PingReply, error)
	grpc.ClientStream
}

type helloServerStreamPingClient struct {
	grpc.ClientStream
}

func (x *helloServerStreamPingClient) Recv() (*PingReply, error) {
	m := new(PingReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloClient) BiDirectionalStreamPing(ctx context.Context, opts ...grpc.CallOption) (Hello_BiDirectionalStreamPingClient, error) {
	stream, err := c.cc.NewStream(ctx, &Hello_ServiceDesc.Streams[2], Hello_BiDirectionalStreamPing_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &helloBiDirectionalStreamPingClient{stream}
	return x, nil
}

type Hello_BiDirectionalStreamPingClient interface {
	Send(*PingRequest) error
	Recv() (*PingReply, error)
	grpc.ClientStream
}

type helloBiDirectionalStreamPingClient struct {
	grpc.ClientStream
}

func (x *helloBiDirectionalStreamPingClient) Send(m *PingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloBiDirectionalStreamPingClient) Recv() (*PingReply, error) {
	m := new(PingReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServer is the server API for Hello service.
// All implementations must embed UnimplementedHelloServer
// for forward compatibility
type HelloServer interface {
	Hello(context.Context, *HelloRequest) (*HelloReply, error)
	ClientStreamPing(Hello_ClientStreamPingServer) error
	ServerStreamPing(*PingRequest, Hello_ServerStreamPingServer) error
	BiDirectionalStreamPing(Hello_BiDirectionalStreamPingServer) error
	mustEmbedUnimplementedHelloServer()
}

// UnimplementedHelloServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServer struct {
}

func (UnimplementedHelloServer) Hello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedHelloServer) ClientStreamPing(Hello_ClientStreamPingServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamPing not implemented")
}
func (UnimplementedHelloServer) ServerStreamPing(*PingRequest, Hello_ServerStreamPingServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamPing not implemented")
}
func (UnimplementedHelloServer) BiDirectionalStreamPing(Hello_BiDirectionalStreamPingServer) error {
	return status.Errorf(codes.Unimplemented, "method BiDirectionalStreamPing not implemented")
}
func (UnimplementedHelloServer) mustEmbedUnimplementedHelloServer() {}

// UnsafeHelloServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServer will
// result in compilation errors.
type UnsafeHelloServer interface {
	mustEmbedUnimplementedHelloServer()
}

func RegisterHelloServer(s grpc.ServiceRegistrar, srv HelloServer) {
	s.RegisterService(&Hello_ServiceDesc, srv)
}

func _Hello_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hello_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hello_ClientStreamPing_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServer).ClientStreamPing(&helloClientStreamPingServer{stream})
}

type Hello_ClientStreamPingServer interface {
	SendAndClose(*PingReply) error
	Recv() (*PingRequest, error)
	grpc.ServerStream
}

type helloClientStreamPingServer struct {
	grpc.ServerStream
}

func (x *helloClientStreamPingServer) SendAndClose(m *PingReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloClientStreamPingServer) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Hello_ServerStreamPing_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServer).ServerStreamPing(m, &helloServerStreamPingServer{stream})
}

type Hello_ServerStreamPingServer interface {
	Send(*PingReply) error
	grpc.ServerStream
}

type helloServerStreamPingServer struct {
	grpc.ServerStream
}

func (x *helloServerStreamPingServer) Send(m *PingReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Hello_BiDirectionalStreamPing_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServer).BiDirectionalStreamPing(&helloBiDirectionalStreamPingServer{stream})
}

type Hello_BiDirectionalStreamPingServer interface {
	Send(*PingReply) error
	Recv() (*PingRequest, error)
	grpc.ServerStream
}

type helloBiDirectionalStreamPingServer struct {
	grpc.ServerStream
}

func (x *helloBiDirectionalStreamPingServer) Send(m *PingReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloBiDirectionalStreamPingServer) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Hello_ServiceDesc is the grpc.ServiceDesc for Hello service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hello_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloService.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Hello_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientStreamPing",
			Handler:       _Hello_ClientStreamPing_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerStreamPing",
			Handler:       _Hello_ServerStreamPing_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BiDirectionalStreamPing",
			Handler:       _Hello_BiDirectionalStreamPing_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/hello.proto",
}
