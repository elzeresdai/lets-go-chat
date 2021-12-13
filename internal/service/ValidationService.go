package service

import (
	"github.com/go-playground/validator/v10"
	"lets-go-chat/internal/models"
)

type ValidationService interface {
	IsValidRequest(user models.UserRequest) bool
}

var ValidationErrs []models.ValidationResponse

func IsValidRequest(user models.UserRequest) bool {
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
