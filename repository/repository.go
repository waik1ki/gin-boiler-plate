package repository

import "gorm.io/gorm"

type Repositories struct {
	User UserRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User: newUserRepository(db),
	}
}
