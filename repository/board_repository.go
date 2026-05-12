package repository

import (
	"time"

	"github.com/adrianus123/project-management/config"
	"github.com/adrianus123/project-management/model"
)

type BoardRepository interface {
	Create(board *model.Board) error
	Update(board *model.Board) error
	FindByPublicID(publicID string) (*model.Board, error)
	AddMember(boardID uint, userIDs []uint) error
}

type boardRepositoryImpl struct {
}

func NewBoardRepository() BoardRepository {
	return &boardRepositoryImpl{}
}

func (r *boardRepositoryImpl) Create(board *model.Board) error {
	return config.DB.Create(board).Error
}

func (r *boardRepositoryImpl) Update(board *model.Board) error {
	return config.DB.Model(&model.Board{}).Where("public_id = ?", board.PublicID).Updates(map[string]interface{}{
		"title":       board.Title,
		"description": board.Description,
		"due_date":    board.DueDate,
	}).Error
}

func (r *boardRepositoryImpl) FindByPublicID(publicID string) (*model.Board, error) {
	var board model.Board
	err := config.DB.Where("public_id = ?", publicID).First(&board).Error
	return &board, err
}

func (r *boardRepositoryImpl) AddMember(boardID uint, userIDs []uint) error {
	if len(userIDs) == 0 {
		return nil
	}

	now := time.Now()
	var members []model.BoardMember

	for _, userID := range userIDs {
		members = append(members, model.BoardMember{
			BoardID:  int64(boardID),
			UserID:   int64(userID),
			JoinedAt: now,
		})
	}

	return config.DB.Create(&members).Error
}
