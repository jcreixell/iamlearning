package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type serviceServer struct {
}

func (s serviceServer) GetMessage(ctx context.Context, message *Message) *Message {
	return message
}

func newServer() *serviceServer {
	s := &serviceServer{}
	return s
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:5555"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		grpcServer := grpc.NewServer()
		RegisterServiceServer(grpcServer, newServer())
		grpcServer.Serve(lis)
	}
}
