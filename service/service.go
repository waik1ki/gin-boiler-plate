package service

import "gin-boiler-plate/repository"

type Services struct {
	User UserService
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		User: newUserService(repos.User),
	}
}
