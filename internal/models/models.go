package models

var UserCollection []User

type User struct {
	Name         string `json:"name"`
	Password     string `json:"password"`
	PasswordHash string `json:"-"`
}
