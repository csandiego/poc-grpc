package main

import (
	server "github.com/csandiego/poc-grpc/go/grpc"
	"github.com/csandiego/poc-grpc/go/protobuf"
	"github.com/csandiego/poc-grpc/go/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	service := service.NewInMemoryBookService()
	server := server.NewBookServiceServer(service)
	grpcServer := grpc.NewServer()
	protobuf.RegisterBookServiceServer(grpcServer, server)
	s, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	grpcServer.Serve(s)
}
