package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DebuggingRoutes(incommingRoutes *gin.Engine, middleware *Middleware) {
	incommingRoutes.GET("/checkhealth", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"result": "",
		})
	})
}
