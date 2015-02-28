package main

import (
	"log"
	"net"

	pb "github.com/yoppi/grpc-example/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type GreeterService struct {
}

const port = ":11111"

func (g *GreeterService) SayHello(c context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {

	reply := &pb.HelloReply{Message: "[" + req.Name + "]" + "hello, world"}
	return reply, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterGreeterServer(server, new(GreeterService))
	server.Serve(lis)
}
