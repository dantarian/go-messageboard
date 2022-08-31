package data

import (
	"pencethren/go-messageboard/entities"
	"pencethren/go-messageboard/repositories"

	"github.com/google/uuid"
)

type inMemoryBoardRepository struct {
	boards []entities.Board
}

func NewInMemoryBoardRepository() repositories.IBoardRepository {
	return &inMemoryBoardRepository{}
}

func (r *inMemoryBoardRepository) Add(board *entities.Board) (uuid.UUID, error) {
	newBoard := *board
	r.boards = append(r.boards, newBoard)
	return newBoard.Id, nil
}

func (r *inMemoryBoardRepository) ExistsWithName(name string) (bool, error) {
	for _, board := range r.boards {
		if board.Name == name {
			return true, nil
		}
	}
	return false, nil
}
