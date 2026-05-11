package repository

import (
	"strings"

	"github.com/adrianus123/project-management/config"
	"github.com/adrianus123/project-management/model"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByID(id uint) (*model.User, error)
	FindByPublicID(publicID string) (*model.User, error)
	FindAllPagination(filter, sort string, limit, offset int) ([]model.User, int64, error)
	Update(user *model.User) error
	Delete(id uint) error
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

func (r *userRepositoryImpl) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := config.DB.First(&user, id).Error
	return &user, err
}

func (r *userRepositoryImpl) FindByPublicID(publicID string) (*model.User, error) {
	var user model.User
	err := config.DB.Where("public_id = ?", publicID).First(&user).Error
	return &user, err
}

func (r *userRepositoryImpl) FindAllPagination(filter, sort string, limit, offset int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	db := config.DB.Model(&model.User{})

	if filter != "" {
		filterPattern := "%" + filter + "%"
		db = db.Where("name Ilike ? OR email Ilike ?", filterPattern, filterPattern)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if sort != "" {
		switch sort {
		case "-id":
			sort = "-internal_id"
		case "id":
			sort = "internal_id"
		}

		if after, ok := strings.CutPrefix(sort, "-"); ok {
			sort = after + " DESC"
		} else {
			sort = sort + " ASC"
		}

		db = db.Order(sort)
	}

	err := db.Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepositoryImpl) Update(user *model.User) error {
	return config.DB.Model(&model.User{}).Where("public_id = ?", user.PublicID).Updates(
		map[string]interface{}{
			"name": user.Name,
		}).Error
}

func (r *userRepositoryImpl) Delete(id uint) error {
	return config.DB.Delete(&model.User{}, id).Error
}
