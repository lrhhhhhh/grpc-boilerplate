package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-boilerplate/api/pb"
	"grpc-boilerplate/interceptor"
	"grpc-boilerplate/internal/config"
	"grpc-boilerplate/internal/logic"
)

func NewGRPCServer(c *config.Config) *grpc.Server {
	creds, err := credentials.NewServerTLSFromFile(c.TLS.CertFile, c.TLS.KeyFile)
	if err != nil {
		panic(err)
	}

	helloInterceptor := new(interceptor.HelloInterceptor)

	grpcServer := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(helloInterceptor.Unary()),
		grpc.StreamInterceptor(helloInterceptor.Stream()),
	)

	helloService := new(logic.HelloService)
	pb.RegisterHelloServer(grpcServer, helloService)

	return grpcServer
}
