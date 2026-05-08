package repository

import (
	"github.com/adrianus123/project-management/config"
	"github.com/adrianus123/project-management/model"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}

type userRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (r *userRepositoryImpl) Create(user *model.User) error {
	return config.DB.Create(user).Error
}

func (r *userRepositoryImpl) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
