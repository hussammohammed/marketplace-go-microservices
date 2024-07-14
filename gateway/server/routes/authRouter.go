package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/root", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "service working")
	})
}
