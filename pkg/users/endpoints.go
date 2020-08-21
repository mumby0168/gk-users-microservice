package users

import (
	"context"

	"github.com/gokit/microservice/pkg/common"

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
		_, err := service.CreateUser(ctx, req.Email, req.FirstName, req.SecondName, req.Password)
		if err != nil {
			return checkForStandardError(err)
		}

		return EmptyResponse{}, nil
	}
}

func checkForStandardError(err error) (interface{}, error) {
	_, ok := err.(common.StandardError)
	if ok {
		return err, nil
	}
	return nil, err
}
