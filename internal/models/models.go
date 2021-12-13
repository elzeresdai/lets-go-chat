package models

import (
	"github.com/google/uuid"
)

// swagger:model UserRequest
type UserRequest struct {
	UserName string `json:"name" validate:"required,min=4"`
	Password string `json:"password" validate:"required,min=8"`
}

// swagger:model CreateUser
type CreateUser struct {
	UserName string
	Hashed   string
}

// swagger:model CreateUserResponse
type CreateUserResponse struct {
	Id       uuid.UUID
	UserName string
}

// swagger:model LoginUserResponse
type LoginUserResponse struct {
	Token string
}
type ValidationResponse struct {
	Message string
}
