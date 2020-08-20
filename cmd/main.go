package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"
	"github.com/gokit/microservice/pkg/users"
	"github.com/gokit/microservice/pkg/users/inmem"
)

func main() {
	router := mux.NewRouter()
	logger := log.NewLogfmtLogger(os.Stderr)
	userRepository := inmem.NewInMemoryUserRepository()
	userService := users.NewUserService(userRepository, logger)

	users.MapUsersPath(userService, router)

	http.ListenAndServe(":8080", router)
}
