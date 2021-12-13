package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
	"lets-go-chat/internal/models"
	"lets-go-chat/internal/service"
	"log"
	"net/http"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(us service.UserService) *UserHandler {
	return &UserHandler{us}
}

//HandleUserCreate function which creating new user
func (h *UserHandler) HandleUserCreate(e echo.Context) error {
	user := models.UserRequest{}
	err := json.NewDecoder(e.Request().Body).Decode(&user)
	defer e.Request().Body.Close()
	if err == io.EOF {
		return e.JSON(http.StatusBadRequest, "Invalid username or password")
	}
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	isValid := service.IsValidRequest(user)
	if isValid {
		newUser, err := h.userService.CreateUser(user)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
		}
		return e.JSON(http.StatusOK, newUser)
	}
	return e.JSON(http.StatusBadRequest, service.ValidationErrs)
}

func (h *UserHandler) HandleUserLogin(e echo.Context) error {
	user := models.UserRequest{}
	err := json.NewDecoder(e.Request().Body).Decode(&user)
	defer e.Request().Body.Close()
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	isValid := service.IsValidRequest(user)
	if isValid {
		logged, err := h.userService.LoginUser(user)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
		}

		if logged == nil {
			return e.JSON(http.StatusBadRequest, "Invalid username or password")
		}

		return e.JSON(http.StatusOK, logged)
	}
	return e.JSON(http.StatusBadRequest, service.ValidationErrs)
}
