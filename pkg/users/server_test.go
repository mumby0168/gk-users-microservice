package users

import (
	"context"
	"net/http"
	"testing"

	"github.com/gorilla/mux"

	"github.com/gokit/microservice/pkg/common"
)

func TestDecodeGetUsersRequest(t *testing.T) {
	//Arrange
	request := &http.Request{}

	//Act
	got, err := decodeGetUsersRequest(context.TODO(), request)

	//Assert
	checkError(err, t)

	if _, ok := got.(*getUsersRequest); !ok {
		t.Errorf("Expected type *users.getUsersRequest but got %T", got)
	}
}

func TestDecodeGetUserRequest_Always_ReturnsIDFromRouteParams(t *testing.T) {
	//Arrange
	id := "2345treghsfdz"
	request := &http.Request{}
	request = mux.SetURLVars(request, map[string]string{"id": id})

	//Act
	got, err := decodeGetUserRequest(context.TODO(), request)

	//Assert
	checkError(err, t)
	if data, ok := got.(*getUserRequest); ok {
		if data.ID != id {
			t.Errorf("Expected id %v but got %v", id, data.ID)
		}
	} else {
		t.Errorf("Expected type *users.getUserRequest but got %T", got)
	}
}

func TestDecodeGetUserRequest_Always_ReturnsErrorWhenNoIdProvided(t *testing.T) {
	//Arrange
	request := &http.Request{}

	//Act
	_, err := decodeGetUserRequest(context.TODO(), request)

	if stderr, ok := err.(common.StandardError); ok {
		if stderr.Code != "bad_params" {
			t.Errorf("Expected code bad_params but got %v", stderr.Code)
		}
	} else {
		t.Errorf("Expected type StandardError but got %T", err)
	}
}

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Did not expect error but got %v", err.Error())
	}
}
