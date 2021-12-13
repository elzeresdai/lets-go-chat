package container

import (
	"lets-go-chat/internal/repository/inMemory"
	"lets-go-chat/internal/service"
)

type Container struct {
	Services Services
}

type Services struct {
	UserService service.UserService
}

// Inject represent the starter of our IoC Container, here we will inject
// the necessary structs/functions that we need to build our project.
func Inject() Container {
	//stores
	userRepo := inMemory.NewUserStoreInMemory()

	//init services
	us := service.NewUserService(userRepo)

	services := Services{
		UserService: us,
	}

	return Container{
		Services: services,
	}
}
