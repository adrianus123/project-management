package service

import (
	"errors"

	"github.com/adrianus123/project-management/model"
	"github.com/adrianus123/project-management/repository"
	"github.com/google/uuid"
)

type BoardService interface {
	Create(board *model.Board) error
	Update(board *model.Board) error
	GetByPublicID(publicID string) (*model.Board, error)
}

type boardServiceImpl struct {
	boardRepository repository.BoardRepository
	userRepository  repository.UserRepository
}

func NewBoardService(boardRepository repository.BoardRepository, userRepository repository.UserRepository) BoardService {
	return &boardServiceImpl{
		boardRepository: boardRepository,
		userRepository:  userRepository,
	}
}

func (s *boardServiceImpl) Create(board *model.Board) error {
	user, err := s.userRepository.FindByPublicID(board.OwnerPublicID.String())
	if err != nil {
		return errors.New("owner not found")
	}

	board.PublicID = uuid.New()
	board.OwnerID = user.InternalID

	return s.boardRepository.Create(board)
}

func (s *boardServiceImpl) Update(board *model.Board) error {
	return s.boardRepository.Update(board)
}

func (s *boardServiceImpl) GetByPublicID(publicID string) (*model.Board, error) {
	return s.boardRepository.FindByPublicID(publicID)
}
