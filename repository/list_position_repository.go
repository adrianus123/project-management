package repository

import (
	"github.com/adrianus123/project-management/config"
	"github.com/adrianus123/project-management/model"
	"github.com/google/uuid"
)

type ListPositionRepository interface {
	GetByBoard(boardPublicID string) (*model.ListPosition, error)
	CreateOrUpdate(boardPublicID string, listOrder []uuid.UUID) error
	GetListOrder(boardPublicID string) ([]uuid.UUID, error)
	UpdateListOrder(position *model.ListPosition) error
}

type listPositionRepositoryImpl struct {
}

func NewListPositionRepository() ListPositionRepository {
	return &listPositionRepositoryImpl{}
}

func (r *listPositionRepositoryImpl) GetByBoard(boardPublicID string) (*model.ListPosition, error) {
	var listPosition model.ListPosition
	err := config.DB.Joins("JOIN boards ON boards.internal_id = list_positions.board_internal_id").
		Where("boards.public_id = ?", boardPublicID).
		First(&listPosition).Error

	return &listPosition, err
}

func (r *listPositionRepositoryImpl) CreateOrUpdate(boardPublicID string, listOrder []uuid.UUID) error {
	return config.DB.Exec(`INSERT INTO list_positions (board_internal_id, list_id)
	SELECT internal_id, ? FROM boards WHERE public_id = ?
	ON CONFLICT (board_internal_id) 
	DO UPDATE SET list_order = EXCLUDE.list_order`, listOrder, boardPublicID).Error
}

func (r *listPositionRepositoryImpl) GetListOrder(boardPublicID string) ([]uuid.UUID, error) {
	pos, err := r.GetByBoard(boardPublicID)
	if err != nil {
		return nil, err
	}
	return pos.ListOrder, nil
}

func (r *listPositionRepositoryImpl) UpdateListOrder(position *model.ListPosition) error {
	return config.DB.Model(&model.ListPosition{}).
		Where("internal_id = ?", position.InternalID).
		Update("list_order", position.ListOrder).Error
}
