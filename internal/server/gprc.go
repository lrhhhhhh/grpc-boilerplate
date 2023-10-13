package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-boilerplate/api/pb"
	"grpc-boilerplate/interceptor"
	"grpc-boilerplate/internal/config"
	"grpc-boilerplate/internal/service"
)

func NewGRPCServer(c *config.Config) (*grpc.Server, error) {
	creds, err := credentials.NewServerTLSFromFile(c.TLS.CertFile, c.TLS.KeyFile)
	if err != nil {
		return nil, err
	}

	helloInterceptor := new(interceptor.HelloInterceptor)

	grpcServer := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(helloInterceptor.Unary()),
		grpc.StreamInterceptor(helloInterceptor.Stream()),
	)

	helloService := new(service.HelloService)
	pb.RegisterHelloServer(grpcServer, helloService)

	return grpcServer, nil
}
