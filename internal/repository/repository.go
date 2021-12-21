package repository

import (
	"lets-go-chat/internal/models"
	modelRepo "lets-go-chat/internal/repository/models"
)

type UserRepository interface {
	LoginUser(user models.LoginUserRequest) (*modelRepo.User, error)
	CreateUser(user models.CreateUserRequest) (*modelRepo.User, error)
}
