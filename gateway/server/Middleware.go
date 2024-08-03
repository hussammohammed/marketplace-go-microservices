package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user "github.com/hussammohammed/marketplace-go-microservices/gateway/user"
)

type Middleware struct {
	userService user.IUserService
}

func NewMiddleware(iUserService user.IUserService) *Middleware {
	return &Middleware{userService: iUserService}
}

func (m *Middleware) AuthAPIRequest(c *gin.Context) {
	authToken := c.Request.Header.Get("Authorization")
	if authToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized user"})
		c.Abort()
	}
	c.Next()
}
