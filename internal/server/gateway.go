package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-boilerplate/api/pb"
	"grpc-boilerplate/internal/config"
	"grpc-boilerplate/internal/service"
)

func NewGatewayMux(c *config.Config) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux()

	if c.Gateway.OnlyUnaryRPC {
		// in-process; only for unary rpc
		helloService := new(service.HelloService)
		if err := pb.RegisterHelloHandlerServer(context.Background(), mux, helloService); err != nil {
			return nil, err
		}
	} else {
		creds, err := credentials.NewClientTLSFromFile(c.TLS.CertFile, c.TLS.ServerNameOverride)
		if err != nil {
			return nil, err
		}
		dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

		// grpc gateway call grpc server
		// for example, grpc server's endpoint=0.0.0.0:9090, grpc gateway listen at 0.0.0.0:9091
		if err := pb.RegisterHelloHandlerFromEndpoint(context.Background(), mux, c.Gateway.Endpoint, dialOptions); err != nil {
			return nil, err
		}
	}
	return mux, nil
}
