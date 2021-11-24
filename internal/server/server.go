package server

import (
	"github.com/labstack/echo/v4"
	"lets-go-chat/internal/handlers"
	"net/http"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/", hello)
	e.POST("/users", handlers.HandleUserCreate)
	e.POST("/login", handlers.HandleUserLogin)
	e.Logger.Fatal(e.Start(":8000"))
	return e
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}
