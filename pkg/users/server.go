package users

import (
	"context"
	"encoding/json"
	"net/http"

	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MapUsersPath(userService UserService, router *mux.Router) {
	userEndpoints := NewUserEndpoints(userService)

	router.Methods("POST").Path("/users").Handler(httpTransport.NewServer(
		userEndpoints.CreateUser,
		decodeCreateUserRequest,
		encodeResponse,
	))
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := &CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
