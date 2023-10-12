package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-boilerplate/api/pb"
	"grpc-boilerplate/etcd"
	"grpc-boilerplate/interceptor"
	"grpc-boilerplate/internal/logic"
	"net"
)

func main() {
	serviceName := "grpc-helloService"
	addr := "0.0.0.0:9090"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	creds, err := credentials.NewServerTLSFromFile("./tls/server.pem", "./tls/t.key")
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

	etcdAddr := "http://localhost:2379"
	etcdConn, err := etcd.New(etcdAddr)
	if err != nil {
		panic(err)
	}

	if err := etcdConn.Register(serviceName, addr); err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
