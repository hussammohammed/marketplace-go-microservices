package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incommingRoutes *gin.Engine, middleware *Middleware) {
	incommingRoutes.GET("/root", middleware.AuthAPIRequest, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "service working")
	})
}
