package models

import (
	"github.com/google/uuid"
)

// swagger:model CreateUserRequest
type CreateUserRequest struct {
	UserName string `json:"userName" validate:"required,min=4"`
	Password string `json:"password" validate:"required,min=8"`
}

// swagger:model LoginUserRequest
type LoginUserRequest struct {
	UserName string `json:"userName" validate:"required,min=4"`
	Password string `json:"password" validate:"required,min=8"`
}

// swagger:model CreateUserResponse
type CreateUserResponse struct {
	Id       uuid.UUID
	UserName string
}

// swagger:model LoginUserResponse
type LoginUserResponse struct {
	Url string
}
type ValidationResponse struct {
	Message string
}
