package main

import (
	"grpcExample/pkg/revers"
	"log"
	"net"

	proto "grpcExample/grpc"

	"google.golang.org/grpc"
)

//go get -u google.golang.org/grpc
//go get -u github.com/golang/protobuf/protoc-gen-go
//go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
//go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

//protoc -I proto --go_out=plugins=grpc:proto proto/proto.proto

func main() {
	log.Print("запуск...")

	s := grpc.NewServer()
	srv := &revers.GRPCServer{}
	proto.RegisterReverseServer(s, srv)
	
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
