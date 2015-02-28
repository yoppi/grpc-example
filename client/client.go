package main

import (
	"log"
	"os"

	pb "github.com/yoppi/grpc-example/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const defaultName = "world"

func main() {
	conn, err := grpc.Dial("127.0.0.1:11111")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.Message)
}
