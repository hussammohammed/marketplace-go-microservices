package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	userService "github.com/hussammohammed/marketplace-go-microservices/gateway/server/grpcClients/protos/user"
)

func DebuggingRoutes(router *gin.Engine, middleware *Middleware, userSvc userService.UserClient) {
	checkhealthGrp := router.Group("/checkhealth")
	checkhealthGrp.GET("/gateway", middleware.AuthAPIRequest, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusOK,
			"result":     "healthy",
		})
	})

	checkhealthGrp.GET("/user", middleware.AuthAPIRequest, func(ctx *gin.Context) {
		if userSvc == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"statusCode": http.StatusInternalServerError,
				"status":     "failed to initialize a connection to user service",
			})
		}
		result, err := userSvc.CheckHealth(ctx, &userService.CheckHealthRequest{})
		if err != nil {
			log.Println(err.Error())
			result = &userService.CheckHealthResponse{StatusCode: http.StatusInternalServerError, Status: "unhealthy"}
		}
		ctx.JSON(http.StatusOK, gin.H{
			"statusCode": result.StatusCode,
			"status":     result.Status,
		})
	})
}
