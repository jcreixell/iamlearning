package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/jcreixell/iamlearning/frameworks/grpc/protocol/protocol"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:5555", "server host:port")
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := protocol.NewServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	msg, error := client.GetMessage(ctx, &protocol.Message{Payload: "Hello"})

	if error != nil {
		log.Fatalf("Error!")
	} else {
		log.Printf("Message: %v", msg.Payload)
	}
}
