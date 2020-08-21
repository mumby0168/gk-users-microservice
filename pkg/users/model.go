package users

import (
	"strings"

	"github.com/gofrs/uuid"
	"github.com/gokit/microservice/pkg/common"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         string
	FirstName  string
	SecondName string
	Email      string
	Hash       string
}

const emptyFieldCode = "empty_field"
const invalidEmailCode = "invalid_email"
const weakPasswordCode = "invalid_password"

func newUser(firstName string, secondName string, email string, password string) (*User, error) {
	if firstName == "" {
		return nil, common.NewStandardError(emptyFieldCode, "First name is required")
	}
	if secondName == "" {
		return nil, common.NewStandardError(emptyFieldCode, "Second name is required")
	}
	if email == "" || !strings.Contains(email, "@") {
		return nil, common.NewStandardError(invalidEmailCode, "The email is invalid")
	}
	if len(password) < 5 {
		return nil, common.NewStandardError(weakPasswordCode, "The password must be more than 5 characters")
	}

	uuid, _ := uuid.NewV4()
	id := uuid.String()

	pwdBytes := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(pwdBytes, bcrypt.MinCost)

	if err != nil {
		return nil, err
	}
	strhash := string(hash)

	return &User{
		ID:         id,
		FirstName:  firstName,
		SecondName: secondName,
		Email:      email,
		Hash:       strhash,
	}, nil
}
