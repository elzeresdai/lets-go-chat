package repository

import (
	"lets-go-chat/internal/models"
	modelRepo "lets-go-chat/internal/repository/models"
)

type UserRepository interface {
	LoginUser(user models.UserRequest) (*modelRepo.User, error)
	CreateUser(user models.CreateUser) (modelRepo.User, error)
}
