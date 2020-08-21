package users

type createUserRequest struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type getUserRequest struct {
	ID string
}

type getUsersRequest struct {
}

type emptyResponse struct {
}

type userResponse struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
}

type usersResponse []userResponse
