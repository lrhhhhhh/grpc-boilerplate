package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

type HelloInterceptor struct {
}

func (i *HelloInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println(info.FullMethod)
		return handler(ctx, req)
	}
}

func (i *HelloInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv any,
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Println(info.FullMethod)
		return handler(srv, ss)
	}
}

func (i *HelloInterceptor) UnaryClient() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		log.Println(method)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
