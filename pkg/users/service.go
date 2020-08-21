package users

import (
	"context"
	"fmt"

	"github.com/gokit/microservice/pkg/common"

	"github.com/go-kit/kit/log"
)

type UserService interface {
	CreateUser(ctx context.Context, email string, firstName string, secondName string, password string) (string, error)
}

type userService struct {
	repository UserRepository
	logger     log.Logger
}

func NewUserService(repository UserRepository, logger log.Logger) UserService {
	return userService{
		repository: repository,
		logger:     logger,
	}
}

func (service userService) CreateUser(ctx context.Context, email string, firstName string, secondName string, password string) (string, error) {

	if service.repository.UserExists(ctx, email) {
		return "", common.NewStandardError("email_in_use", "The email supplied is in use")
	}

	user, err := newUser(firstName, secondName, email, password)
	if err != nil {
		_ = fmt.Errorf("User creation failed: %v", err.Error())
		return "", err
	}

	err = service.repository.CreateUser(ctx, user)

	if err != nil {
		return "", nil
	}

	fmt.Println("Created user with id:", user.ID)
	return user.ID, nil

}
