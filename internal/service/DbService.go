package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"lets-go-chat/internal/models"
	"lets-go-chat/internal/repository"
	"os"
	"time"
)

type userServiceImplDB struct {
	userRepoDB repository.UserRepoDB
}

func NewUserServiceDB(user repository.UserRepoDB) UserService {
	return userServiceImplDB{user}
}

func (us userServiceImplDB) CreateUser(user models.CreateUserRequest) (*models.CreateUserResponse, error) {
	newUser, err := us.userRepoDB.CreateUser(user)

	if err != nil {
		return nil, err
	}
	resp := &models.CreateUserResponse{
		Id:       newUser.ID,
		UserName: newUser.Name,
	}
	return resp, nil
}

func (us userServiceImplDB) LoginUser(user models.LoginUserRequest) (*models.LoginUserResponse, error) {
	existUser, _ := us.userRepoDB.LoginUser(user)
	token, err := us.GenerateToken(existUser.ID)
	if err != nil {
		return nil, err
	}
	resp := models.LoginUserResponse{Url: token}
	return &resp, nil

}

func (us userServiceImplDB) GenerateToken(userId uuid.UUID) (string, error) {
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
