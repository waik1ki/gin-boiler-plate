package service

import (
	"gin-boiler-plate/model"
	"gin-boiler-plate/repository"
)

type UserService interface {
	CreateUser(user *model.User) error
	GetUsers() ([]model.User, error)
	GetUserByID(id uint) (model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func newUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(user *model.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetUsers() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s *userService) GetUserByID(id uint) (model.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) UpdateUser(user *model.User) error {
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
