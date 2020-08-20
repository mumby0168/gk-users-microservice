package users

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type UserEndpoints struct {
	CreateUser endpoint.Endpoint
}

type CreateUserRequest struct {
	FirstName  string `json:="firstName"`
	SecondName string `json:="secondName"`
	Email      string `json:="email"`
	Password   string `json:="password"`
}

type EmptyResponse struct {
}

func NewUserEndpoints(service UserService) *UserEndpoints {
	return &UserEndpoints{
		CreateUser: createUserEndpoint(service),
	}
}

func createUserEndpoint(service UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*CreateUserRequest)
		_, err := service.CreateUser(ctx, req.FirstName, req.SecondName, req.Email, req.Password)
		return EmptyResponse{}, err
	}
}
