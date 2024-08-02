package server

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	userService "github.com/hussammohammed/marketplace-go-microservices/gateway/server/grpcClients/protos/user"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	userSvcAddr = flag.String("userSvcAddr", "localhost:50051", "the address to connect to")
)

func Run() error {
	flag.Parse()
	// create gin router with default middleware
	router := gin.Default()
	// CORS config
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "Access-Control-Allow-Origin"}
	router.Use(cors.New(config))

	// create middleware
	middleware := NewMiddleware()
	// set all system routes
	getRoutes(router, middleware)
	return router.Run(fmt.Sprintf("%v:%v", viper.GetString("server.host"), viper.GetString("server.port")))
}

func getRoutes(router *gin.Engine, middleware *Middleware) {
	AuthRoutes(router, middleware)
	// initialize grpc services
	// 1- initialize user service
	userSvcClient := initUserSvcConnection()
	DebuggingRoutes(router, middleware, userSvcClient)
}

func initUserSvcConnection() userService.UserClient {
	var userSvcClient userService.UserClient
	userSvcConn, err := grpc.Dial(*userSvcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("faild to connect to user service: %v", err)
	} else {
		defer userSvcConn.Close()
		userSvcClient = userService.NewUserClient(userSvcConn)
	}
	return userSvcClient
}
