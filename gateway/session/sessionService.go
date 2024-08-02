package session

import (
	user "github.com/hussammohammed/marketplace-go-microservices/gateway/user"
)

type ISessionService interface {
	AuthUser(loginData user.LoginDto) (string, error)
	ValidateAuthToken(token string) (bool, error)
}

type SessionService struct {
	userService user.IUserService
}

func NewSessionService(iUserService user.IUserService) *SessionService {
	return &SessionService{userService: iUserService}
}

func (s *SessionService) AuthUser(loginData user.LoginDto) (string, error) {
	return "abc", nil
}
