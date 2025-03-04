package handler

import "gin-boiler-plate/service"

type Handlers struct {
	User *UserHandler
}

func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		User: newUserHandler(services.User),
	}
}
