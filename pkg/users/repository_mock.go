package users

import (
	"context"
)

type userRepositoryMock struct {
	CreateUserFunc func(context.Context, *User) error
	UserExistsFunc func(context.Context, string) bool
	GetUserFunc    func(context.Context, string) (User, bool)
	GetUsersFunc   func(context.Context) (*[]User, error)
}

func (m userRepositoryMock) CreateUser(ctx context.Context, user *User) error {
	return m.CreateUserFunc(ctx, user)
}

func (m userRepositoryMock) GetUsers(ctx context.Context) (*[]User, error) {
	return m.GetUsersFunc(ctx)
}

func (m userRepositoryMock) UserExists(ctx context.Context, email string) bool {
	return m.UserExistsFunc(ctx, email)
}

func (m userRepositoryMock) GetUser(ctx context.Context, ID string) (User, error) {
	return m.GetUser(ctx, ID)
}
