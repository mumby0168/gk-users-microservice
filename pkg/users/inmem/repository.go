package inmem

import (
	"context"

	"github.com/gokit/microservice/pkg/users"
)

type inMemoryUserRepository struct {
}

func NewInMemoryUserRepository() users.UserRepository {
	return inMemoryUserRepository{}
}

func (r inMemoryUserRepository) CreateUser(ctx context.Context, user users.User) error {
	return nil
}
