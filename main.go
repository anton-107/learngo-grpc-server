package main

import (
	"flag"
	"github.com/anton-107/learngo-grpc-server/proto"
	"github.com/gengo/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

const rpcAddress = ":50051"
const restAddress = ":3001"

func startRpcServer() {
	listener, err := net.Listen("tcp", rpcAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	helloworld.RegisterGreeterServer(grpcServer, &server{})
	go grpcServer.Serve(listener)

	log.Println("Started grpc service at " + rpcAddress)

}

func startRestServer() {
	echoEndpoint := flag.String("greeter_endpoint", rpcAddress, "greeter endpoint")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	err := helloworld.RegisterGreeterHandlerFromEndpoint(ctx, mux, *echoEndpoint)
	if err != nil {
		if err != nil {
			log.Fatalf("failed to register a handler: %v", err)
		}
	}

	http.ListenAndServe(restAddress, mux)
	log.Println("Started rest service at " + restAddress)
}

func main() {
	log.Println("Start RPC")
	startRpcServer()
	log.Println("Start Rest")
	startRestServer()

}
