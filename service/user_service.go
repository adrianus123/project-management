package service

import (
	"errors"

	"github.com/adrianus123/project-management/constant"
	"github.com/adrianus123/project-management/model"
	"github.com/adrianus123/project-management/repository"
	"github.com/adrianus123/project-management/util"
	"github.com/google/uuid"
)

type UserService interface {
	Register(user *model.User) error
	Login(email, password string) (*model.User, error)
	GetUserByID(id uint) (*model.User, error)
	GetUserByPublicID(publicID string) (*model.User, error)
	GetAllPagination(filter, sort string, limit, offset int) ([]model.User, int64, error)
	Update(user *model.User) error
	Delete(id string) error
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
		return errors.New(constant.ERR_EMAIL_ALREADY_EXISTS)
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

func (s *userServiceImpl) Login(email, password string) (*model.User, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return nil, errors.New(constant.ERR_INVALID_CREDENTIAL)
	}

	if !util.VerifyPassword(password, user.Password) {
		return nil, errors.New(constant.ERR_INVALID_CREDENTIAL)
	}

	return user, nil
}

func (s *userServiceImpl) GetUserByID(id uint) (*model.User, error) {
	return s.userRepository.FindByID(id)
}

func (s *userServiceImpl) GetUserByPublicID(publicID string) (*model.User, error) {
	return s.userRepository.FindByPublicID(publicID)
}

func (s *userServiceImpl) GetAllPagination(filter, sort string, limit, offset int) ([]model.User, int64, error) {
	return s.userRepository.FindAllPagination(filter, sort, limit, offset)
}

func (s *userServiceImpl) Update(user *model.User) error {
	return s.userRepository.Update(user)
}

func (s *userServiceImpl) Delete(id string) error {
	user, err := s.userRepository.FindByPublicID(id)
	if err != nil {
		return errors.New(err.Error())
	}

	if user == nil {
		return errors.New(constant.ERR_DATA_NOT_FOUND)
	}

	return s.userRepository.Delete(uint(user.InternalID))
}
