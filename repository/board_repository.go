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
	RemoveMembers(boardID uint, userIDs []uint) error
	FindAllByUserPaginate(userPublicID, filter, sort string, limit, offset int) ([]model.Board, int64, error)
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

func (r *boardRepositoryImpl) RemoveMembers(boardID uint, userIDs []uint) error {
	if len(userIDs) == 0 {
		return nil
	}

	return config.DB.Where("board_internal_id = ? AND user_internal_id IN (?)", boardID, userIDs).
		Delete(&model.BoardMember{}).Error
}

func (r *boardRepositoryImpl) FindAllByUserPaginate(userPublicID, filter, sort string, limit, offset int) ([]model.Board, int64, error) {
	var boards []model.Board
	var total int64

	query := config.DB.Model(&model.Board{}).
		Where("owner_public_id = ? OR internal_id IN ("+
			"SELECT bm.board_internal_id FROM board_members bm JOIN users u ON u.internal_id = bm.user_internal_id "+
			"WHERE u.public_id = ?)", userPublicID, userPublicID)

	// Filter by title
	if filter != "" {
		query = query.Where("title ILIKE ?", "%"+filter+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Sorting by created date
	if sort != "" {
		query = query.Order(sort)
	} else {
		query = query.Order("created_at desc")
	}

	if err := query.Limit(limit).Offset(offset).Find(&boards).Error; err != nil {
		return nil, 0, err
	}

	return boards, total, nil
}
