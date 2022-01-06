package service

import (
	"github.com/google/uuid"
	"lets-go-chat/internal/models"
)

type UserService interface {
	CreateUser(user models.CreateUserRequest) (*models.CreateUserResponse, error)
	LoginUser(user models.LoginUserRequest) (*models.LoginUserResponse, error)
	GenerateToken(userId uuid.UUID) (string, error)
}
