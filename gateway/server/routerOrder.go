package server

import (
	"github.com/gin-gonic/gin"
	order "github.com/hussammohammed/marketplace-go-microservices/gateway/servicesHandlers/order"
)

func OrderRoutes(router *gin.Engine, middleware *Middleware, orderController *order.OrderController) {
	userGrp := router.Group("/order")
	userGrp.POST("/new", middleware.AuthAPIRequest, orderController.CreateOrder)
}
