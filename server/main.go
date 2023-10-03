package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"grpc-demo/interceptor"
	pb "grpc-demo/pb"
	"log"
	"math/rand"
	"net"
	"time"
)

type HelloService struct {
	pb.UnimplementedHelloServer
}

func (s *HelloService) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println(req.Msg)
	return &pb.HelloReply{Content: "hello from the other side ~"}, nil
}

func (s *HelloService) ClientStreamPing(stream pb.Hello_ClientStreamPingServer) error {
	numPing := new(pb.NumPing)
	if err := stream.RecvMsg(numPing); err != nil {
		log.Println(err)
		return err
	}

	for i := 1; i <= int(numPing.Cnt); i++ {
		pingReq, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("RecvClientStreamPing, total=%d, now=%d, timestamp=%v\n", numPing.Cnt, i, pingReq.Timestamp)
	}
	return nil
}

func (s *HelloService) ServerStreamPing(req *pb.PingRequest, stream pb.Hello_ServerStreamPingServer) error {
	var err error
	rand.Seed(time.Now().Unix())
	n := rand.Intn(10) + 1
	if err := stream.SendMsg(&pb.NumPing{Cnt: int64(n)}); err != nil {
		return err
	}

	ticker := time.NewTicker(time.Millisecond * 500)
	for i := 1; i <= n; i++ {
		select {
		case <-ticker.C:
			if err = stream.SendMsg(&pb.PingRequest{Timestamp: time.Now().Unix()}); err != nil {
				log.Println(err)
				break
			} else {
				log.Printf("ServerStreamSendPing, total=%d, now=%d, timestamp=%d\n", n, i, req.Timestamp)
			}
		}
	}
	return err
}

func (s *HelloService) BiDirectionalStreamPing(stream pb.Hello_BiDirectionalStreamPingServer) error {
	var err error
	ch := make(chan *pb.PingRequest)

	// reader
	go func() {
		for {
			if req, err := stream.Recv(); err != nil {
				log.Println(err)
				if status.Code(err) == codes.Canceled {
					break
				} else {
					continue
				}
			} else {
				ch <- req
				log.Printf("RecvBiDirectionalPing, timestamp=%v\n", req.Timestamp)
			}
		}
	}()

	// writer
	for {
		select {
		case <-ch:
			if err = stream.Send(&pb.PingReply{Timestamp: time.Now().Unix()}); err != nil {
				log.Println(err)
				if status.Code(err) == codes.Canceled {
					break
				}
			}
		}
	}
	return err
}

func main() {
	listen, _ := net.Listen("tcp", ":9090")
	creds := insecure.NewCredentials()

	helloInterceptor := new(interceptor.HelloInterceptor)

	grpcServer := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(helloInterceptor.Unary()),
		grpc.StreamInterceptor(helloInterceptor.Stream()),
	)

	helloService := new(HelloService)
	pb.RegisterHelloServer(grpcServer, helloService)

	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
