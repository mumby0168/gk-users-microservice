package users

type User struct {
	ID         string
	FirstName  string
	SecondName string
	Email      string
	Hash       []byte
	Salt       []byte
}
