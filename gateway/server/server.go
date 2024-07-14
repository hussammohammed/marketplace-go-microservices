package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	routes "github.com/hussammohammed/marketplace-go-microservices/gateway/server/routes"
	"github.com/spf13/viper"
)

func Run() error {
	// create gin router with default middleware
	router := gin.Default()
	// CORS config
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "Access-Control-Allow-Origin"}
	router.Use(cors.New(config))

	routes.DebuggingRoutes(router)
	routes.AuthRoutes(router)
	return router.Run(fmt.Sprintf("%v:%v", viper.GetString("server.host"), viper.GetString("server.port")))
}
