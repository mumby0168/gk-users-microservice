package users

import (
	"context"
	"fmt"

	"github.com/gokit/microservice/pkg/common"

	"github.com/go-kit/kit/log"
)

type UserService interface {
	CreateUser(ctx context.Context, email string, firstName string, secondName string, password string) (string, error)
	GetUser(ctx context.Context, ID string) (*User, error)
	GetUsers(ctx context.Context) ([]User, error)
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

func (s userService) CreateUser(ctx context.Context, email string, firstName string, secondName string, password string) (string, error) {

	if s.repository.UserExists(ctx, email) {
		return "", common.NewStandardError("email_in_use", "The email supplied is in use")
	}

	user, err := newUser(firstName, secondName, email, password)
	if err != nil {
		_ = fmt.Errorf("User creation failed: %v", err.Error())
		return "", err
	}

	err = s.repository.CreateUser(ctx, user)

	if err != nil {
		return "", nil
	}

	fmt.Println("Created user with id:", user.ID)
	return user.ID, nil

}

func (s userService) GetUser(ctx context.Context, ID string) (*User, error) {
	user, err := s.repository.GetUser(ctx, ID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s userService) GetUsers(ctx context.Context) ([]User, error) {
	users, err := s.repository.GetUsers(ctx)

	if err != nil {
		return nil, err
	}

	return *users, nil
}
