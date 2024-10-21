package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/hussammohammed/marketplace-go-microservices/microservices/user/config"
	"github.com/hussammohammed/marketplace-go-microservices/microservices/user/db"
	protos "github.com/hussammohammed/marketplace-go-microservices/microservices/user/protos"
	"github.com/hussammohammed/marketplace-go-microservices/microservices/user/repository"
	"github.com/hussammohammed/marketplace-go-microservices/microservices/user/server"
	userModule "github.com/hussammohammed/marketplace-go-microservices/microservices/user/userModule"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	initConfig()
	// DB
	accountsDb, err := db.Dial(viper.GetString("db.url"))
	if err != nil {
		log.Fatalf("failed to open db connection: %v", err)
		return
	}
	defer accountsDb.Close()
	// register repositories
	userRepo := repository.NewUserRepository(accountsDb)
	// register services
	userSvc := userModule.NewUserService(userRepo)
	u := server.NewUser(userSvc)
	// init grpc server
	grpcServer := grpc.NewServer()
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

func initConfig() {
	err := config.Load("development", strings.Split("./config", ",")...)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("config file used:%v", viper.ConfigFileUsed())
}
