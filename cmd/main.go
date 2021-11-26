package main

import (
	"lets-go-chat/internal/server"
)

func main() {
	e := server.New()
	e.Logger.Fatal(e.Start(":8000"))
	return
}
