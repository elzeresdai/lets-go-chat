package handlers

import "lets-go-chat/internal/container"

type handlers struct {
	User *UserHandler
}

func InitHandlers(di container.Container) handlers {
	return handlers{
		User: NewUserHandler(di.Services.UserService),
	}
}
