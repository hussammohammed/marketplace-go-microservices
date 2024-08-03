package server

import (
	"github.com/gin-gonic/gin"
	user "github.com/hussammohammed/marketplace-go-microservices/gateway/user"
)

func UserRoutes(router *gin.Engine, middleware *Middleware, UserController *user.UserController) {
	userGrp := router.Group("/user")
	userGrp.POST("/login", UserController.Login)
}
