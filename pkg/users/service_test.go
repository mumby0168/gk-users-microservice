package users

import (
	"context"
	"os"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestCreateUser_Always_CreatesUser(t *testing.T) {
	//Arrange
	mock := &userRepositoryMock{}
	logger := log.NewLogfmtLogger(os.Stderr)
	sut := NewUserService(mock, logger)
	createCalled := false
	checkedExisting := false

	mock.UserExistsFunc = func(ctx context.Context, email string) bool {
		checkedExisting = true
		return false
	}

	mock.CreateUserFunc = func(ctx context.Context, user *User) error {
		createCalled = true
		return nil
	}

	//Act

	_, err := sut.CreateUser(context.TODO(), "test@test.com", "billy", "mumby", "Test123456")

	//Assert
	if err != nil {
		t.Errorf("Expected to no error but failed with: %v", err.Error())
	}

	if !createCalled {
		t.Errorf("Expected save user but no call was made")
	}

	if !checkedExisting {
		t.Errorf("Expected check for existing")
	}
}
