package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	userService "github.com/hussammohammed/marketplace-go-microservices/gateway/server/grpcClients/protos/user"
)

func DebuggingRoutes(incommingRoutes *gin.Engine, userSvc userService.UserClient) {
	incommingRoutes.GET("/checkhealth", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"result": "",
		})
	})

	incommingRoutes.GET("/user/checkhealth", func(ctx *gin.Context) {
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
