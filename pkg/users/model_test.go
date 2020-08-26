package users

import (
	"testing"

	"github.com/gokit/microservice/pkg/common"
)

func Test_NewUser_Always_CreatesValidUser(t *testing.T) {
	//Arrange
	fname := "Joe"
	sname := "Bloggs"
	email := "test@test.com"
	password := "Test123456"

	//Act
	user, _ := newUser(fname, sname, email, password)

	//Assert
	if user == nil {
		t.Errorf("User should not be nil")
		return
	}

	if user.FirstName != fname {
		t.Errorf("expected first name %v and got %v", fname, user.FirstName)
		return
	}

	if user.SecondName != sname {
		t.Errorf("expected second name %v and got %v", sname, user.SecondName)
		return
	}

	if user.Email != email {
		t.Errorf("expected email %v and got %v", email, user.Email)
		return
	}

	if user.Hash == "" {
		t.Errorf("expected hash to have value but was empty")
		return
	}

	if user.ID == "" {
		t.Errorf("Expected id to contain data but is empty")
	}
}

func Test_NewUser_Always_ReturnsErrorWhenFirstNameNotSupplied(t *testing.T) {
	//Arrange
	fname := ""
	sname := "Bloggs"
	email := "test@test.com"
	password := "Test123456"

	//Act
	_, err := newUser(fname, sname, email, password)

	if err, ok := err.(common.StandardError); !ok || err.Code != emptyFieldCode {
		t.Errorf("expected common.StandardError with code: %v but got %v", emptyFieldCode, err.Error())
	}
}

func Test_NewUser_Always_ReturnsErrorWhenSecondNameNotSupplied(t *testing.T) {
	//Arrange
	fname := "Joe"
	sname := ""
	email := "test@test.com"
	password := "Test123456"

	//Act
	_, err := newUser(fname, sname, email, password)

	//Assert
	if err, ok := err.(common.StandardError); !ok || err.Code != emptyFieldCode {
		t.Errorf("expected common.StandardError with code: %v but got %v", emptyFieldCode, err.Error())
	}
}

func Test_NewUser_Always_ReturnsErrorWhenEmailInvalid(t *testing.T) {
	//Arrange
	fname := "Joe"
	sname := "Bloggs"
	password := "Test123456"

	emails := [2]string{"", "bademailaddress.com"}

	for _, email := range emails {
		//Act
		_, err := newUser(fname, sname, email, password)

		//Assert
		if err, ok := err.(common.StandardError); !ok || err.Code != invalidEmailCode {
			t.Errorf("expected common.StandardError with code: %v but got %v", invalidEmailCode, err.Error())
		}
	}

}

func Test_NewUser_Always_ReturnsErrorWhenPasswordInvalid(t *testing.T) {
	//Arrange
	fname := "Joe"
	sname := "Bloggs"
	email := "test@test.com"
	password := "Test"

	//Act
	_, err := newUser(fname, sname, email, password)

	//Assert
	if err, ok := err.(common.StandardError); !ok || err.Code != weakPasswordCode {
		t.Errorf("expected common.StandardError with code: %v but got %v", weakPasswordCode, err.Error())
	}
}
