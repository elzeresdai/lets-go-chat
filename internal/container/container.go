package container

import (
	postgresdb2 "lets-go-chat/internal/db"
	"lets-go-chat/internal/repository/postgres"
	"lets-go-chat/internal/service"
)

type Container struct {
	Services Services
}

type Services struct {
	UserService service.UserService
}

func Inject() Container {
	db := postgresdb2.ConnectDB()
	userRepo := postgres.NewUserRepo(db)
	us := service.NewUserServiceDB(userRepo)

	services := Services{
		UserService: us,
	}

	return Container{
		Services: services,
	}
}
