package inMemory

import (
	"github.com/google/uuid"
	"lets-go-chat/internal/models"
	"lets-go-chat/internal/repository"
	modelsRepo "lets-go-chat/internal/repository/models"
	"lets-go-chat/pkg/hasher"
	"sync"
)

type userStoreInMemory struct {
	users []modelsRepo.User
	s     sync.RWMutex
}

func NewUserStoreInMemory() repository.UserRepository {
	return &userStoreInMemory{
		users: make([]modelsRepo.User, 0),
	}
}

func (repo *userStoreInMemory) CreateUser(user models.CreateUserRequest) (*modelsRepo.User, error) {
	repo.s.Lock()
	defer repo.s.Unlock()

	hash, err := hasher.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	newUser := modelsRepo.User{
		ID:     uuid.New(),
		Name:   user.UserName,
		Hashed: hash,
	}
	repo.users = append(repo.users, newUser)
	return &newUser, nil

}

func (repo *userStoreInMemory) LoginUser(user models.LoginUserRequest) (*modelsRepo.User, error) {
	repo.s.Lock()
	defer repo.s.Unlock()
	for _, i := range repo.users {
		if i.Name == user.UserName && hasher.CheckPasswordHash(user.Password, i.Hashed) {
			return &i, nil
		}
	}
	return nil, nil
}
