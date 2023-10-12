package main

import (
	"grpc-boilerplate/etcd"
	"grpc-boilerplate/internal/config"
	"grpc-boilerplate/internal/server"
	"net"
)

func main() {
	var cfg config.Config
	cfg.Service.Name = "grpc-helloService"
	cfg.Service.Addr = "0.0.0.0:9090"
	cfg.TLS.CertFile = "./tls/server.pem"
	cfg.TLS.KeyFile = "./tls/t.key"
	cfg.Etcd.Addr = "http://localhost:2379"

	listen, err := net.Listen("tcp", cfg.Service.Addr)
	if err != nil {
		panic(err)
	}

	etcdConn, err := etcd.New(cfg.Etcd.Addr)
	if err != nil {
		panic(err)
	}

	if err := etcdConn.Register(cfg.Service.Name, cfg.Service.Addr); err != nil {
		panic(err)
	}

	grpcServer := server.NewGRPCServer(&cfg)
	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
