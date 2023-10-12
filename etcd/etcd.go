package etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	gresolver "google.golang.org/grpc/resolver"
	"log"
	"time"
)

type Conn struct {
	addr   string
	client *clientv3.Client
}

func New(addr string) (*Conn, error) {
	cfg := clientv3.Config{
		Endpoints:            []string{addr},
		DialTimeout:          10 * time.Second,
		DialKeepAliveTime:    10 * time.Second,
		DialKeepAliveTimeout: 10 * time.Second,
	}
	client, err := clientv3.New(cfg)
	if err != nil {
		return nil, err
	}

	return &Conn{
		addr:   addr,
		client: client,
	}, nil
}

func (c *Conn) Register(serviceName, serverAddr string) error {
	mgr, err := endpoints.NewManager(c.client, serviceName)
	if err != nil {
		log.Println(err)
		return err
	}

	lease, err := c.client.Grant(context.Background(), 10)
	if err != nil {
		log.Println(err)
		return err
	}

	if err = mgr.AddEndpoint(
		context.Background(),
		fmt.Sprintf("%s/%s", serviceName, serverAddr),
		endpoints.Endpoint{Addr: serverAddr}, clientv3.WithLease(lease.ID),
	); err != nil {
		return err
	}

	alive, err := c.client.KeepAlive(context.Background(), lease.ID)
	if err != nil {
		return err
	}

	go func() {
		for {
			<-alive
			log.Println("etcd server keep alive")
		}
	}()
	return nil
}

func (c *Conn) UnRegister(serviceName, serverAddr string) error {
	if c.client != nil {
		mgr, err := endpoints.NewManager(c.client, serviceName)
		if err != nil {
			return err
		}
		if err = mgr.DeleteEndpoint(
			context.TODO(),
			fmt.Sprintf("%s/%s", serviceName, serverAddr),
		); err != nil {
			return err
		}
		return err
	}
	return nil
}

func (c *Conn) Resolver() (gresolver.Builder, error) {
	return resolver.NewBuilder(c.client)
}
