package main

import (
	"grpc-boilerplate/internal/config"
	"grpc-boilerplate/internal/server"
	"net"
	"net/http"
)

func main() {
	var cfg config.Config
	cfg.Gateway.Enable = true
	cfg.Gateway.Addr = "0.0.0.0:9091"
	cfg.Gateway.Endpoint = "0.0.0.0:9090"
	cfg.Gateway.OnlyUnaryRPC = false
	cfg.TLS.Enable = true
	cfg.TLS.CertFile = "./tls/server.pem"
	cfg.TLS.KeyFile = "./tls/server.key"
	cfg.TLS.ServerNameOverride = "*.example.com"

	listener, err := net.Listen("tcp", cfg.Gateway.Addr)
	if err != nil {
		panic(err)
	}

	mux, err := server.NewGatewayMux(&cfg)
	if err != nil {
		panic(err)
	}

	if cfg.Gateway.TLS.Enable {
		// another cert and key for https
		if err := http.ServeTLS(listener, mux, cfg.Gateway.TLS.CertFile, cfg.Gateway.TLS.KeyFile); err != nil {
			panic(err)
		}
	} else {
		if err := http.Serve(listener, mux); err != nil {
			panic(err)
		}
	}
}
