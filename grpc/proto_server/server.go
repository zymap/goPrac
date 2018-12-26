package main

import (
	pb "../proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"fmt"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *pb.RequestArgs) (*pb.ResponseArgs, error) {
	return &pb.ResponseArgs{Message: string(in.Message)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":10999")
	fmt.Println(lis.Addr().String())
	if err != nil {
		log.Fatalf("error", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("error", err)
	}
}
