package main

import (
	"flag"
	"lets-go-chat/internal/server"
	"os"
)

type options struct {
	port string
}

var opt options

func init() {
	flag.StringVar(&opt.port, "p", os.Getenv("PORT"), "The default port to listen on")
	flag.Parse()

	if opt.port == "" {
		opt.port = "7"
	}
}

func main() {
	e := server.New()
	e.Logger.Fatal(e.Start(":8000"))
	return
}
