package server

import (
	"context"
	"net/http"

	protos "github.com/hussammohammed/marketplace-go-microservices/microservices/user/protos"
	userModule "github.com/hussammohammed/marketplace-go-microservices/microservices/user/userModule"
)

type User struct {
	protos.UnimplementedUserServer
	userService userModule.IuserService
}

func NewUser(iUserSvc userModule.IuserService) *User {
	return &User{userService: iUserSvc}
}

func (s *User) CheckHealth(ctx context.Context, rr *protos.CheckHealthRequest) (*protos.CheckHealthResponse, error) {
	return &protos.CheckHealthResponse{StatusCode: http.StatusOK, Status: "user service is healthy"}, nil
}
func (s *User) CreateUser(ctx context.Context, request *protos.CreateUserRequest) (*protos.CreateUserResponse, error) {
	err := s.userService.CreateUser(request)
	if err != nil {
		return &protos.CreateUserResponse{StatusCode: http.StatusInternalServerError, Message: err.Error()}, err
	}
	return &protos.CreateUserResponse{StatusCode: http.StatusOK, Message: "user created successfully"}, nil
}
