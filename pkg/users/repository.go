package users

import (
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	UserExists(ctx context.Context, email string) bool
}
