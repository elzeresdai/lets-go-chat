package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"lets-go-chat/internal/models"
	"lets-go-chat/pkg/hasher"
	"log"
	"net/http"
)

//HandleUserCreate function which creating new user
func HandleUserCreate(e echo.Context) error {
	user := models.User{}
	err := json.NewDecoder(e.Request().Body).Decode(&user)
	defer e.Request().Body.Close()
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	user.PasswordHash, err = hasher.HashPassword(user.Password)
	if err != nil {
		return err
	}
	models.UserCollection = append(models.UserCollection, user)
	log.Printf("this is new user %#v, password hash %#v", user.Name, user.PasswordHash)
	return e.String(http.StatusOK, "User created!")
}
