package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"lets-go-chat/internal/models"
	"lets-go-chat/pkg/hasher"
	"log"
	"net/http"
)

//HandleUserLogin function checking users name and password
func HandleUserLogin(e echo.Context) error {
	user := models.User{}
	err := json.NewDecoder(e.Request().Body).Decode(&user)
	defer e.Request().Body.Close()
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	user.PasswordHash, err = hasher.HashPassword(user.Password)
	if err != nil {
		return e.String(http.StatusBadRequest, "Something went wrong check your password")
	}
	for _, i := range models.UserCollection {
		if i.Name == user.Name && hasher.CheckPasswordHash(user.Password, i.PasswordHash) {
			return e.String(http.StatusOK, "You are successfully logged in")
		}
	}

	return e.String(http.StatusUnauthorized, "Check your name or password")
}
