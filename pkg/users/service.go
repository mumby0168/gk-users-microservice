package users

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"

	"github.com/gofrs/uuid"
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
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	return id, errors.New("Something went wrong")
}
