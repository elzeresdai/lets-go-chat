package service

import (
	"github.com/go-playground/validator/v10"
	"lets-go-chat/internal/models"
)

type ValidationService interface {
	IsValidRequest(user models.CreateUserRequest) bool
	IsValidLoginRequest(user models.LoginUserRequest) bool
}

var ValidationErrs []models.ValidationResponse

func IsValidRequest(user models.CreateUserRequest) bool {
	v := validator.New()
	err := v.Struct(user)

	if err == nil {
		return true
	}
	for _, er := range err.(validator.ValidationErrors) {
		errors := models.ValidationResponse{
			Message: er.Error(),
		}
		ValidationErrs = append(ValidationErrs, errors)
	}

	return false

}

func IsValidLoginRequest(user models.LoginUserRequest) bool {
	v := validator.New()
	err := v.Struct(user)

	if err == nil {
		return true
	}
	for _, er := range err.(validator.ValidationErrors) {
		errors := models.ValidationResponse{
			Message: er.Error(),
		}
		ValidationErrs = append(ValidationErrs, errors)
	}

	return false

}
