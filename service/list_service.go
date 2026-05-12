package service

import (
	"errors"
	"fmt"

	"github.com/adrianus123/project-management/config"
	"github.com/adrianus123/project-management/model"
	"github.com/adrianus123/project-management/model/types"
	"github.com/adrianus123/project-management/repository"
	"github.com/adrianus123/project-management/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ListWithOrder struct {
	Positions []uuid.UUID
	Lists     []model.List
}

type ListService interface {
	GetByBoardID(boardPublicID string) (*ListWithOrder, error)
	GetByID(id uint) (*model.List, error)
	GetByPublicID(publicID string) (*model.List, error)
	Create(list *model.List) error
	Update(list *model.List) error
	Delete(id uint) error
	UpdatePosition(boardPublicID string, positions []uuid.UUID) error
}

type listServiceImpl struct {
	listRepository         repository.ListRepository
	boardRepository        repository.BoardRepository
	listPositionRepository repository.ListPositionRepository
}

func NewListService(listRepository repository.ListRepository, boardRepository repository.BoardRepository, listPositionRepository repository.ListPositionRepository) ListService {
	return &listServiceImpl{
		listRepository:         listRepository,
		boardRepository:        boardRepository,
		listPositionRepository: listPositionRepository,
	}
}

func (s *listServiceImpl) GetByBoardID(boardPublicID string) (*ListWithOrder, error) {
	_, err := s.boardRepository.FindByPublicID(boardPublicID)
	if err != nil {
		return nil, errors.New("board not found")
	}

	positions, err := s.listPositionRepository.GetListOrder(boardPublicID)
	if err != nil {
		return nil, errors.New("failed to get list order : " + err.Error())
	}

	if len(positions) == 0 {
		return nil, errors.New("list position not found")
	}

	lists, err := s.listRepository.FindByBoardID(boardPublicID)
	if err != nil {
		return nil, errors.New("failed to get list : " + err.Error())
	}

	orderedList := util.SortListByPosition(lists, positions)

	return &ListWithOrder{
		Positions: positions,
		Lists:     orderedList,
	}, nil
}

func (s *listServiceImpl) GetByID(id uint) (*model.List, error) {
	return s.listRepository.FindByID(id)
}

func (s *listServiceImpl) GetByPublicID(publicID string) (*model.List, error) {
	return s.listRepository.FindByPublicID(publicID)
}

func (s *listServiceImpl) Create(list *model.List) error {
	board, err := s.boardRepository.FindByPublicID(list.BoardPublicID.String())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("board not found")
		}
		return fmt.Errorf("failed to find board : %s", err.Error())
	}

	list.BoardInternalID = board.InternalID

	if list.PublicID == uuid.Nil {
		list.PublicID = uuid.New()
	}

	// start transaction
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(list).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create list : %w", err)
	}

	var position model.ListPosition
	res := tx.Where("board_internal_id = ?", board.InternalID).First(&position)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		// create new if not exist
		position = model.ListPosition{
			PublicID:  uuid.New(),
			BoardID:   board.InternalID,
			ListOrder: types.UUIDArray{list.PublicID},
		}

		if err := tx.Create(&position).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to create list position : %w", err)
		}
	} else if res.Error != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create list position : %w", res.Error)
	} else {
		position.ListOrder = append(position.ListOrder, list.PublicID)

		if err := tx.Model(&position).Update("list_order", position.ListOrder).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update list position : %w", err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction : %w", err)
	}

	return nil
}

func (s *listServiceImpl) Update(list *model.List) error {
	return s.listRepository.Update(list)
}

func (s *listServiceImpl) Delete(id uint) error {
	return s.listRepository.Delete(id)
}

func (s *listServiceImpl) UpdatePosition(boardPublicID string, positions []uuid.UUID) error {
	board, err := s.boardRepository.FindByPublicID(boardPublicID)
	if err != nil {
		return errors.New("board not found")
	}

	position, err := s.listPositionRepository.GetByBoard(board.PublicID.String())
	if err != nil {
		return errors.New("list position not found")
	}

	position.ListOrder = positions

	return s.listPositionRepository.UpdateListOrder(position)
}
