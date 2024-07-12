package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	routes "github.com/hussammohammed/marketplace-go-microservices/gateway/routes"
)

func main() {
	fmt.Println("starting user service")
	// create gin router with default middleware
	router := gin.Default()
	// CORS config
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "Access-Control-Allow-Origin"}
	router.Use(cors.New(config))

	routes.AuthRoutes(router)
	router.GET("/checkhealth", checkHealthStatus)
	router.Run(":8080")
}

func checkHealthStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "healthy",
		"result": "",
	})
}
