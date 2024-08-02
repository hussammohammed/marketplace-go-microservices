package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) AuthAPIRequest(c *gin.Context) {
	authToken := c.Request.Header.Get("Authorization")
	if authToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized user"})
		c.Abort()
	}
	c.Next()
}
