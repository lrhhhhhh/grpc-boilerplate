package main

import (
	"context"
	pb "grpc-boilerplate/api/pb"
	"grpc-boilerplate/internal/client"
	"grpc-boilerplate/internal/config"
	"log"
	"math/rand"
	"time"
)

func unary(client pb.HelloClient) {
	if reply, err := client.Hello(context.Background(), &pb.HelloRequest{Msg: &pb.Message{
		FromUsername: "WangXiaoMei",
		ToUsername:   "WangXiaoGang",
		Content:      "Excuse me?",
	}}); err != nil {
		panic(err)
	} else {
		log.Println(reply)
	}
}

func clientStream(client pb.HelloClient) {
	if stream, err := client.ClientStreamPing(context.Background()); err != nil {
		panic(err)
	} else {
		rand.Seed(time.Now().Unix())
		n := rand.Intn(10) + 1
		if err := stream.SendMsg(&pb.NumPing{Cnt: int64(n)}); err != nil {
			panic(err)
		}
		ticker := time.NewTicker(time.Millisecond * 500)
		for i := 1; i <= n; i++ {
			select {
			case <-ticker.C:
				if err := stream.SendMsg(&pb.PingRequest{Timestamp: time.Now().Unix()}); err != nil {
					panic(err)
				} else {
					log.Printf("ClientStreamSendPing, total=%d, now=%d\n", n, i)
				}
			}
		}
		if err := stream.CloseSend(); err != nil {
			panic(err)
		}
		ticker.Stop()
	}
}

func serverStream(client pb.HelloClient) {
	if stream, err := client.ServerStreamPing(context.Background(), &pb.PingRequest{Timestamp: time.Now().Unix()}); err != nil {
		panic(err)
	} else {
		numPing := new(pb.NumPing)
		if err := stream.RecvMsg(numPing); err != nil {
			panic(err)
		}
		for i := 1; i <= int(numPing.Cnt); i++ {
			if reply, err := stream.Recv(); err != nil {
				log.Println(err)
				break
			} else {
				log.Printf("RecvServerStreamPing, total=%d, now=%d, timestamp=%v\n", numPing.Cnt, i, reply.Timestamp)
			}
		}
	}
}

func biDirectionalStream(client pb.HelloClient) {
	if stream, err := client.BiDirectionalStreamPing(context.Background()); err != nil {
		panic(err)
	} else {
		// reader
		go func() {
			ticker := time.NewTicker(time.Millisecond * 500)
			for {
				select {
				case <-ticker.C:
					if reply, err := stream.Recv(); err != nil {
						log.Println(err, "?")
						continue
					} else {
						log.Printf("ServerRecvBiDirectionalPingStream, ts=%v\n", reply.Timestamp)
					}
				}
			}
		}()

		// writer
		ticker := time.NewTicker(time.Millisecond * 500)
		for {
			select {
			case <-ticker.C:
				nts := time.Now().Unix()
				if err := stream.Send(&pb.PingRequest{Timestamp: nts}); err != nil {
					log.Println(err, "??")
					continue
				} else {
					log.Printf("ServerSendBiDirectionalPingStream, ts=%v\n", nts)
				}
			}
		}
	}
}

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

	conn, err := client.NewGRPCClient(&cfg)
	if err != nil {
		panic(err)
	}

	grpcClient := pb.NewHelloClient(conn)

	unary(grpcClient)
	clientStream(grpcClient)
	serverStream(grpcClient)
	biDirectionalStream(grpcClient)
}
