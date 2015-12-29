package main

import (
	"flag"
	"github.com/anton-107/learngo-grpc-server/proto"
	"github.com/gengo/grpc-gateway/runtime"
	"github.com/hashicorp/consul/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

const protocol = "http"  // should be reachable from Consul agent
const host = "localhost" // should be reachable from Consul agent and API gateway
const rpcAddress = ":50051"
const restPort = 3005

var restPortStr = strconv.Itoa(restPort)
var restAddress = ":" + restPortStr

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

	log.Println("Started rest service at " + restAddress)
	http.ListenAndServe(restAddress, mux)
}

func registerInServiceDiscovery() {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	err = client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      strings.Join([]string{"helloworld", host, restPortStr}, "-"),
		Name:    "/v1/helloworld/say",
		Address: host,
		Port:    restPort,
		Tags:    []string{"go", "microservice", "api", "rest"},
		Check: &api.AgentServiceCheck{
			HTTP:     protocol + "://" + host + ":" + restPortStr + "/healthcheck",
			Interval: "10s",
		},
	})

	if err != nil {
		log.Fatalf("Error registering service for discovery: %v", err)
	}
}

func main() {
	log.Println("Start RPC")
	startRpcServer()

	log.Println("Register for service discovery")
	registerInServiceDiscovery()

	log.Println("Start Rest")
	startRestServer()

	for {
	}
}
