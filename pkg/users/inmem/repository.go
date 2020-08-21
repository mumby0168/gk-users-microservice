package inmem

import (
	"context"

	"github.com/gokit/microservice/pkg/users"
)

type inMemoryUserRepository struct {
	users map[string]*users.User
}

func NewInMemoryUserRepository() users.UserRepository {
	return inMemoryUserRepository{
		users: make(map[string]*users.User),
	}
}

func (r inMemoryUserRepository) CreateUser(ctx context.Context, user *users.User) error {
	r.users[user.ID] = user
	return nil
}

func (r inMemoryUserRepository) UserExists(ctx context.Context, email string) bool {

	for _, user := range r.users {
		if user.Email == email {
			return true
		}
	}

	return false
}
