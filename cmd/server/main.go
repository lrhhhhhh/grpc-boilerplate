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
	cfg.TLS.Enable = true
	cfg.TLS.CertFile = "./tls/server.pem"
	cfg.TLS.KeyFile = "./tls/server.key"
	cfg.TLS.ServerNameOverride = "*.example.com"
	cfg.Discovery.Enable = true
	cfg.Discovery.Etcd.Addr = "http://localhost:2379"

	listen, err := net.Listen("tcp", cfg.Service.Addr)
	if err != nil {
		panic(err)
	}

	if cfg.Discovery.Enable {
		etcdConn, err := etcd.New(cfg.Discovery.Etcd.Addr)
		if err != nil {
			panic(err)
		}

		if err := etcdConn.Register(cfg.Service.Name, cfg.Service.Addr); err != nil {
			panic(err)
		}
	}

	grpcServer, err := server.NewGRPCServer(&cfg)
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
