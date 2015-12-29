package main

import (
	"github.com/anton-107/learngo-grpc-server/proto"
	"golang.org/x/net/context"
)

type server struct{}

func (server *server) Healthcheck(context context.Context, msg *helloworld.HealthRequest) (*helloworld.HealthResponse, error) {
	return &helloworld.HealthResponse{
		Status: "OK",
	}, nil
}

func (server *server) SayWord(context context.Context, msg *helloworld.MyMessage) (*helloworld.MyMessage, error) {
	msg.Text = "Hello " + msg.Text
	return msg, nil
}
