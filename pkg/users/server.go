package users

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gokit/microservice/pkg/common"

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

	router.Methods("GET").Path("/users/{id}").Handler(httpTransport.NewServer(
		userEndpoints.GetUser,
		decodeGetUserRequest,
		encodeResponse,
	))

	router.Methods("GET").Path("/users").Handler(httpTransport.NewServer(
		userEndpoints.GetUsers,
		decodeGetUsersRequest,
		encodeResponse,
	))
}

func decodeGetUsersRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return &getUsersRequest{}, nil
}

func decodeGetUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	if id, ok := vars["id"]; ok {
		return &getUserRequest{ID: id}, nil
	}

	return nil, common.NewStandardError("bad_params", "The id parameter must be provided by the route")
}

func decodeCreateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := &createUserRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	errored, err := common.ErrorChecks(w, response)

	if err != nil {
		return err
	}

	if errored {
		return nil
	}

	return common.WriteJson(w, response)
}
