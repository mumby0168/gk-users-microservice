package inmem

import (
	"context"

	"github.com/gokit/microservice/pkg/common"

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

func (r inMemoryUserRepository) GetUser(ctx context.Context, ID string) (users.User, error) {
	if val, ok := r.users[ID]; ok {
		return *val, nil
	}
	return users.User{}, common.NewNotFoundError("A user with identifier: " + ID + " is not in our records")
}

func (r inMemoryUserRepository) GetUsers(ctx context.Context) (*[]users.User, error) {
	usrs := make([]users.User, 0, len(r.users))
	for _, user := range r.users {
		usrs = append(usrs, *user)
	}
	return &usrs, nil
}
