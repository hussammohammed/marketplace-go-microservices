package server

import (
	"context"
	"net/http"

	protos "github.com/hussammohammed/marketplace-go-microservices/microservices/user/protos"
)

type User struct {
	protos.UnimplementedUserServer
}

func NewUser() *User {
	return &User{}
}

func (s *User) CheckHealth(ctx context.Context, rr *protos.CheckHealthRequest) (*protos.CheckHealthResponse, error) {
	return &protos.CheckHealthResponse{StatusCode: http.StatusOK, Status: "user service is healthy"}, nil
}
