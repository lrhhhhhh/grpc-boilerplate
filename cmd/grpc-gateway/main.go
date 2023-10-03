package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-demo/pb"
	"grpc-demo/service"
	"net"
	"net/http"
)

const onlyUnaryRPC = false
const enableTLS = false

func main() {
	addr := "0.0.0.0:9091"
	endpoint := "0.0.0.0:9090"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()

	if onlyUnaryRPC {
		// in-process
		// only for unary rpc
		helloService := new(service.HelloService)
		if err := pb.RegisterHelloHandlerServer(ctx, mux, helloService); err != nil {
			panic(err)
		}
	} else {
		creds, err := credentials.NewClientTLSFromFile("./tls/server.pem", "*.liuronghao.com")
		if err != nil {
			panic(err)
		}
		dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

		// grpc gateway call grpc server
		// for example, grpc server's endpoint=0.0.0.0:9090, grpc gateway listen at 0.0.0.0:9091
		if err := pb.RegisterHelloHandlerFromEndpoint(ctx, mux, endpoint, dialOptions); err != nil {
			panic(err)
		}
	}

	if enableTLS {
		// another cert and key for https
		certPath := ""
		keyPath := ""
		if err := http.ServeTLS(listener, mux, certPath, keyPath); err != nil {
			panic(err)
		}
	} else {
		if err := http.Serve(listener, mux); err != nil {
			panic(err)
		}
	}
}
