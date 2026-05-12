package repository

import (
	"github.com/adrianus123/project-management/config"
	"github.com/adrianus123/project-management/model"
)

type BoardMemberRepository interface {
	GetMembers(boardPublicID string) ([]model.User, error)
}

type boardMemberRepositoryImpl struct {
}

func NewBoardMemberRepository() BoardMemberRepository {
	return &boardMemberRepositoryImpl{}
}

func (r *boardMemberRepositoryImpl) GetMembers(boardPublicID string) ([]model.User, error) {
	var users []model.User

	err := config.DB.Joins("JOIN board_members bm ON bm.user_internal_id = users.internal_id").
		Joins("JOIN boards b ON b.internal_id = bm.board_internal_id").
		Where("b.public_id = ?", boardPublicID).
		Find(&users).Error

	return users, err
}
