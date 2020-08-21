package users

import (
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	UserExists(ctx context.Context, email string) bool
	GetUser(ctx context.Context, ID string) (User, error)
	GetUsers(ctx context.Context) (*[]User, error)
}
