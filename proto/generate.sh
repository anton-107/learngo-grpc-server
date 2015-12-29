#!/usr/bin/env bash

# proto:
#protoc -I ./ ./helloworld.proto --go_out=plugins=grpc:./

# Generate gRPC stub:
protoc -I/usr/local/include -I. \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/gengo/grpc-gateway/third_party/googleapis \
 --go_out=Mgoogle/api/annotations.proto=github.com/gengo/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
 ./helloworld.proto

# Generate reverse-proxy:
protoc -I/usr/local/include -I. \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/gengo/grpc-gateway/third_party/googleapis \
 --grpc-gateway_out=logtostderr=true:. \
 ./helloworld.proto
