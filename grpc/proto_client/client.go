package main

import (
	"google.golang.org/grpc"
	"log"
	pb "../proto"
	"context"
	"time"
)

const (
	address     = "127.0.0.1:10999"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect")
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	name := defaultName

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.RequestArgs{Message: name})
	if err != nil {
		log.Fatalf("error", err)
	}

	log.Println("message : ", r.Message)

}
