package data

import (
	"pencethren/go-messageboard/entity"
	"pencethren/go-messageboard/repository"

	"github.com/google/uuid"
)

type inMemoryBoardRepository struct {
	boards []entity.Board
}

func NewInMemoryBoardRepository() repository.IBoardRepository {
	return &inMemoryBoardRepository{}
}

func (r *inMemoryBoardRepository) Add(board *entity.Board) (uuid.UUID, error) {
	newBoard := *board
	r.boards = append(r.boards, newBoard)
	return newBoard.Id, nil
}

func (r *inMemoryBoardRepository) List(pageSize int, filter *entity.BoardSearch) ([]*entity.BoardSummary, error) {
	boards := []*entity.BoardSummary{}
	for _, board := range r.boards {
		boards = append(boards, &entity.BoardSummary{
			Id:    board.Id,
			Name:  board.Name,
			State: board.State,
		})
	}
	return boards, nil
}

func (r *inMemoryBoardRepository) ExistsWithName(name string) (bool, error) {
	for _, board := range r.boards {
		if board.Name == name {
			return true, nil
		}
	}
	return false, nil
}
