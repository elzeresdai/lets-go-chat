package service

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/google/uuid"
	"lets-go-chat/internal/models"
	"lets-go-chat/internal/repository"
	"os"
	"time"
)

type UserService interface {
	CreateUser(user models.CreateUserRequest) (*models.CreateUserResponse, error)
	LoginUser(user models.LoginUserRequest) (*models.LoginUserResponse, error)
	GenerateToken(userId uuid.UUID) (string, error)
}

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(user repository.UserRepository) UserService {
	return userServiceImpl{user}
}

func (us userServiceImpl) CreateUser(user models.CreateUserRequest) (*models.CreateUserResponse, error) {
	newUser, err := us.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	resp := &models.CreateUserResponse{
		Id:       newUser.ID,
		UserName: newUser.Name,
	}
	return resp, nil
}

func (us userServiceImpl) LoginUser(user models.LoginUserRequest) (*models.LoginUserResponse, error) {
	existUser, _ := us.userRepo.LoginUser(user)
	if existUser == nil {
		return nil, nil
	}

	token, err := us.GenerateToken(existUser.ID)
	if err != nil {
		return nil, err
	}
	resp := models.LoginUserResponse{Url: token}
	return &resp, nil
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
