package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	protos "github.com/hussammohammed/marketplace-go-microservices/microservices/user/protos"
	"github.com/hussammohammed/marketplace-go-microservices/microservices/user/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	grpcServer := grpc.NewServer()
	u := server.NewUser()
	protos.RegisterUserServer(grpcServer, u)
	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(grpcServer)
	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	log.Printf("server listing at: %v", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
