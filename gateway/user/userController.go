package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	Login(ctx *gin.Context)
}

type UserController struct {
	userService IUserService
}

func NewUserController(iUserService IUserService) *UserController {
	return &UserController{userService: iUserService}
}
func (uc *UserController) Login(ctx *gin.Context) {
	var loginDto LoginDto
	err := ctx.Bind(&loginDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token := uc.userService.Login(&loginDto)
	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userName or password"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
