package user

import (
	userMicroService "github.com/hussammohammed/marketplace-go-microservices/gateway/server/grpcClients/protos/user"
)

type IUserService interface {
	Login(loginDto *LoginDto) string
	ValidateAuthToken(token string) (bool, error)
}
type UserService struct {
	userMicroService *userMicroService.UserClient
}

func NewUserService(userMicroSvc *userMicroService.UserClient) *UserService {
	return &UserService{userMicroService: userMicroSvc}
}
func (u *UserService) Login(loginDto *LoginDto) string {
	if loginDto.Email == "hos" {
		return "abd"
	}
	return ""
}

func (u *UserService) ValidateAuthToken(token string) (bool, error) {
	return false, nil
}
