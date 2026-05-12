package repository

import (
	"github.com/adrianus123/project-management/config"
	"github.com/adrianus123/project-management/model"
	"github.com/google/uuid"
)

type ListRepository interface {
	Create(list *model.List) error
	Update(list *model.List) error
	Delete(id uint) error
	UpdatePosition(boardPublicID string, positions []string) error
	GetCardPosition(listPublicID string) ([]uuid.UUID, error)
	FindByBoardID(boardPublicID string) ([]model.List, error)
	FindByPublicID(publicID string) (*model.List, error)
	FindByID(id uint) (*model.List, error)
}

type listRepositoryImpl struct{}

func NewListRepository() ListRepository {
	return &listRepositoryImpl{}
}

func (r *listRepositoryImpl) Create(list *model.List) error {
	return config.DB.Create(list).Error
}

func (r *listRepositoryImpl) Update(list *model.List) error {
	return config.DB.Model(&model.List{}).Where("public_id = ?", list.PublicID).Updates(map[string]interface{}{
		"title": list.Title,
	}).Error
}

func (r *listRepositoryImpl) Delete(id uint) error {
	return config.DB.Delete(&model.List{}, id).Error
}

func (r *listRepositoryImpl) UpdatePosition(boardPublicID string, positions []string) error {
	return config.DB.Model(&model.ListPosition{}).Where("board_internal_id = (SELECT internal_id FROM boards WHERE public_id = ?)", boardPublicID).Update("list_order", positions).Error
}

func (r *listRepositoryImpl) GetCardPosition(listPublicID string) ([]uuid.UUID, error) {
	var positions model.CardPosition

	err := config.DB.Joins("JOIN lists ON lists.internal_id = card_positions.list_internal_id").Where("lists.public_id = ?", listPublicID).Find(&positions).Error

	return positions.CardOrder, err
}

func (r *listRepositoryImpl) FindByBoardID(boardPublicID string) ([]model.List, error) {
	var lists []model.List

	err := config.DB.Where("board_public_id = ?", boardPublicID).
		Order("internal_id ASC").
		Find(&lists).Error

	return lists, err
}

func (r *listRepositoryImpl) FindByPublicID(publicID string) (*model.List, error) {
	var list model.List

	err := config.DB.Where("public_id = ?", publicID).First(&list).Error

	return &list, err
}

func (r *listRepositoryImpl) FindByID(id uint) (*model.List, error) {
	var list model.List

	err := config.DB.First(&list, id).Error

	return &list, err
}
