package client

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-boilerplate/etcd"
	"grpc-boilerplate/interceptor"
	"grpc-boilerplate/internal/config"
	"log"
)

func NewGRPCClient(c *config.Config) (conn *grpc.ClientConn, err error) {
	var creds credentials.TransportCredentials
	if c.TLS.Enable {
		creds, err = credentials.NewClientTLSFromFile(c.TLS.CertFile, c.TLS.ServerNameOverride)
		if err != nil {
			return nil, err
		}
	} else {
		creds = insecure.NewCredentials()
	}

	// interceptor
	helloInterceptor := new(interceptor.HelloInterceptor)

	if c.Discovery.Enable {
		etcdConn, err := etcd.New(c.Discovery.Etcd.Addr)
		if err != nil {
			panic(err)
		}

		resolver, err := etcdConn.Resolver()
		if err != nil {
			panic(err)
		}
		conn, err = grpc.Dial(
			fmt.Sprintf("etcd:///%s", c.Service.Name),
			grpc.WithResolvers(resolver),
			grpc.WithTransportCredentials(creds),
			grpc.WithUnaryInterceptor(helloInterceptor.UnaryClient()),
		)
	} else {
		conn, err = grpc.Dial(
			c.Service.Addr,
			grpc.WithTransportCredentials(creds),
		)
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return conn, nil
}
