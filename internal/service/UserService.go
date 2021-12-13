package service

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/google/uuid"
	"lets-go-chat/internal/models"
	"lets-go-chat/internal/repository"
	"lets-go-chat/pkg/hasher"
	"os"
	"time"
)

type UserService interface {
	CreateUser(user models.UserRequest) (*models.CreateUserResponse, error)
	LoginUser(user models.UserRequest) (*models.LoginUserResponse, error)
	HashPassword(password string) (string, error)
	GenerateToken(userId uuid.UUID) (string, error)
}

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(user repository.UserRepository) UserService {
	return userServiceImpl{user}
}

func (us userServiceImpl) CreateUser(user models.UserRequest) (*models.CreateUserResponse, error) {
	hashed, err := us.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	userRequest := models.CreateUser{
		UserName: user.UserName,
		Hashed:   hashed,
	}
	newUser, err := us.userRepo.CreateUser(userRequest)
	if err != nil {
		return nil, err
	}
	resp := &models.CreateUserResponse{
		Id:       newUser.ID,
		UserName: newUser.Name,
	}
	return resp, nil
}

func (us userServiceImpl) LoginUser(user models.UserRequest) (*models.LoginUserResponse, error) {
	existUser, _ := us.userRepo.LoginUser(user)
	if existUser == nil {
		return nil, nil
	}

	token, err := us.GenerateToken(existUser.ID)
	if err != nil {
		return nil, err
	}
	resp := models.LoginUserResponse{Token: token}
	return &resp, nil
}

func (us userServiceImpl) HashPassword(password string) (string, error) {
	hash, err := hasher.HashPassword(password)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (us userServiceImpl) GenerateToken(userId uuid.UUID) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
