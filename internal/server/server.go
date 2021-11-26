package server

import (
	"github.com/labstack/echo/v4"
	"lets-go-chat/internal/handlers"
	"log"
	"net/http"
	"os"
)

func New() *echo.Echo {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	e := echo.New()
	e.GET("/", hello)
	e.POST("/users", handlers.HandleUserCreate)
	e.POST("/login", handlers.HandleUserLogin)
	e.Logger.Fatal(e.Start(":" + port))
	return e
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}
