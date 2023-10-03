package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-demo/etcd"
	"grpc-demo/interceptor"
	pb "grpc-demo/pb"
	"grpc-demo/service"
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

	helloService := new(service.HelloService)
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
