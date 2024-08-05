package server

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hussammohammed/marketplace-go-microservices/gateway/helpers"
	userMicroService "github.com/hussammohammed/marketplace-go-microservices/gateway/server/grpcClients/protos/user"
	user "github.com/hussammohammed/marketplace-go-microservices/gateway/user"
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
	// initialize grpc services
	// 1- initialize user service
	userSvcClient := initUserSvcConnection()

	//helpers
	cryptHelper := helpers.NewCryptHelper()

	// services
	userSvc := user.NewUserService(userSvcClient, cryptHelper)

	// controllers
	userCtrl := user.NewUserController(userSvc)
	// create middleware
	middleware := NewMiddleware(userSvc)
	// set all system routes
	UserRoutes(router, middleware, userCtrl)
	DebuggingRoutes(router, middleware, userSvcClient)
	return router.Run(fmt.Sprintf("%v:%v", viper.GetString("server.host"), viper.GetString("server.port")))
}

func initUserSvcConnection() userMicroService.UserClient {
	var userSvcClient userMicroService.UserClient
	userSvcConn, err := grpc.Dial(*userSvcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("faild to connect to user service: %v", err)
	} else {
		defer userSvcConn.Close()
		userSvcClient = userMicroService.NewUserClient(userSvcConn)
	}
	return userSvcClient
}
