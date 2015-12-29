package main

import (
	"github.com/anton-107/learngo-grpc-server/proto"
	"golang.org/x/net/context"
)

type server struct{}

func (server *server) SayWord(context context.Context, msg *helloworld.MyMessage) (*helloworld.MyMessage, error) {
	msg.Text = "Hello " + msg.Text
	return msg, nil
}
