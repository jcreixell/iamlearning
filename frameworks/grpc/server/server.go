package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jcreixell/grpc/protocol"
	"google.golang.org/grpc"
)

type serviceServer struct {
}

func (s serviceServer) GetMessage(ctx context.Context, message *protocol.Message) (*protocol.Message, error) {
	return &protocol.Message{Payload: "Echo: " + message.Payload}, nil
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
		protocol.RegisterServiceServer(grpcServer, newServer())
		grpcServer.Serve(lis)
	}
}
