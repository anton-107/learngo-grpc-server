package main

import (
	"github.com/anton-107/learngo-grpc-server/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	helloworld.RegisterGreeterServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}
