package service

import (
	"errors"

	"github.com/adrianus123/project-management/model"
	"github.com/adrianus123/project-management/repository"
	"github.com/adrianus123/project-management/util"
	"github.com/google/uuid"
)

type UserService interface {
	Register(user *model.User) error
}

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{userRepository: userRepository}
}

func (s *userServiceImpl) Register(user *model.User) error {
	userExist, _ := s.userRepository.FindByEmail(user.Email)
	if userExist.InternalID != 0 {
		return errors.New("Email already registered")
	}

	hashed, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashed
	user.Role = "user"
	user.PublicID = uuid.New()

	return s.userRepository.Create(user)
}
