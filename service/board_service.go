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
	AddMembers(boardPublicID string, userPublicIDs []string) error
	RemoveMembers(boardPublicID string, userPublicIDs []string) error
	GetAllByUserPaginate(userPublicID, filter, sort string, limit, offset int) ([]model.Board, int64, error)
}

type boardServiceImpl struct {
	boardRepository       repository.BoardRepository
	userRepository        repository.UserRepository
	boardMemberRepository repository.BoardMemberRepository
}

func NewBoardService(boardRepository repository.BoardRepository, userRepository repository.UserRepository, boardMemberRepository repository.BoardMemberRepository) BoardService {
	return &boardServiceImpl{
		boardRepository:       boardRepository,
		userRepository:        userRepository,
		boardMemberRepository: boardMemberRepository,
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

func (s *boardServiceImpl) AddMembers(boardPublicID string, userPublicIDs []string) error {
	board, err := s.boardRepository.FindByPublicID(boardPublicID)
	if err != nil {
		return errors.New("board not found")
	}

	var userIDs []uint
	for _, userPublicID := range userPublicIDs {
		user, err := s.userRepository.FindByPublicID(userPublicID)
		if err != nil {
			return errors.New("user not found")
		}
		userIDs = append(userIDs, uint(user.InternalID))
	}

	members, err := s.boardMemberRepository.GetMembers(boardPublicID)
	if err != nil {
		return err
	}

	memberMap := make(map[uint]bool)
	for _, member := range members {
		memberMap[uint(member.InternalID)] = true
	}

	var newMemberIDs []uint
	for _, userID := range userIDs {
		if _, exists := memberMap[userID]; !exists {
			newMemberIDs = append(newMemberIDs, userID)
		}
	}

	if len(newMemberIDs) == 0 {
		return nil
	}

	return s.boardRepository.AddMember(uint(board.InternalID), newMemberIDs)
}

func (s *boardServiceImpl) RemoveMembers(boardPublicID string, userPublicIDs []string) error {
	board, err := s.boardRepository.FindByPublicID(boardPublicID)
	if err != nil {
		return errors.New("board not found")
	}

	var userIDs []uint
	for _, userPublicID := range userPublicIDs {
		user, err := s.userRepository.FindByPublicID(userPublicID)
		if err != nil {
			return errors.New("user not found")
		}
		userIDs = append(userIDs, uint(user.InternalID))
	}

	members, err := s.boardMemberRepository.GetMembers(boardPublicID)
	if err != nil {
		return err
	}

	memberMap := make(map[uint]bool)
	for _, member := range members {
		memberMap[uint(member.InternalID)] = true
	}

	var removeMemberIDs []uint
	for _, userID := range userIDs {
		if _, exists := memberMap[userID]; exists {
			removeMemberIDs = append(removeMemberIDs, userID)
		}
	}

	if len(removeMemberIDs) == 0 {
		return nil
	}

	return s.boardRepository.RemoveMembers(uint(board.InternalID), removeMemberIDs)
}

func (s *boardServiceImpl) GetAllByUserPaginate(userPublicID, filter, sort string, limit, offset int) ([]model.Board, int64, error) {
	return s.boardRepository.FindAllByUserPaginate(userPublicID, filter, sort, limit, offset)
}
