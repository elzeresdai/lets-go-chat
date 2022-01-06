package postgres

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"lets-go-chat/internal/models"
	"lets-go-chat/internal/repository"
	modelsRepo "lets-go-chat/internal/repository/models"
	"lets-go-chat/pkg/hasher"
	"log"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) repository.UserRepoDB {
	return &UserRepo{
		db: db,
	}
}

func (u UserRepo) LoginUser(user models.LoginUserRequest) (*modelsRepo.UserDB, error) {
	query := `SELECT id, hash FROM users WHERE name = $1 AND hash = $2`
	hashed, _ := hasher.HashPassword(user.Password)
	var existUser modelsRepo.UserDB
	err := u.db.Get(&existUser, query, user.UserName, hashed)
	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return nil, err
	}
	return &existUser, nil
}

func (u UserRepo) CreateUser(user models.CreateUserRequest) (*modelsRepo.UserDB, error) {
	hash, err := hasher.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	query := `INSERT INTO users ( id , name , hash) VALUES ($1, $2, $3) RETURNING $1, $2 `
	var newResp modelsRepo.UserDB
	err = u.db.QueryRow(query, uuid.New(), user.UserName, hash).Scan(&newResp.ID, &newResp.Name)
	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return nil, err
	}
	return &newResp, nil
}
