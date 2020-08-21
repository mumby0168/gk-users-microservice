package users

import (
	"context"

	"github.com/gokit/microservice/pkg/common"

	"github.com/go-kit/kit/endpoint"
)

type UserEndpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
	GetUsers   endpoint.Endpoint
}

func NewUserEndpoints(service UserService) *UserEndpoints {
	return &UserEndpoints{
		CreateUser: createUserEndpoint(service),
		GetUser:    createGetUserEndpoint(service),
		GetUsers:   createGetUsersEndpoint(service),
	}
}

func createGetUserEndpoint(service UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*getUserRequest)
		user, err := service.GetUser(ctx, req.ID)

		if err != nil {
			return checkForServiceErrors(err)
		}

		return userResponse{
			FirstName:  user.FirstName,
			SecondName: user.SecondName,
			Email:      user.Email,
		}, nil
	}
}

func createGetUsersEndpoint(service UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(*getUsersRequest)
		users, err := service.GetUsers(ctx)
		if err != nil {
			return checkForServiceErrors(err)
		}

		var usersDtos []userResponse

		for _, user := range users {
			usersDtos = append(usersDtos, userResponse{
				FirstName:  user.FirstName,
				SecondName: user.SecondName,
				Email:      user.Email,
			})
		}

		return usersDtos, nil
	}
}

func createUserEndpoint(service UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*createUserRequest)
		_, err := service.CreateUser(ctx, req.Email, req.FirstName, req.SecondName, req.Password)
		if err != nil {
			return checkForServiceErrors(err)
		}

		return emptyResponse{}, nil
	}
}

func checkForServiceErrors(err error) (interface{}, error) {
	switch t := err.(type) {
	case common.NotFoundError:
		return t, nil
	case common.StandardError:
		return t, nil
	default:
		return nil, err
	}
}
