package data

import (
	"fmt"
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

	if len(r.boards) == 0 {
		return boards, nil
	}

	start := 0

	if filter.Bookmark != "" {
		found := false
		for idx, board := range r.boards {
			if board.Name == filter.Bookmark {
				start = idx
				found = true
				break
			}
		}

		if !found {
			return boards, fmt.Errorf("bookmark not found: %v", filter.Bookmark)
		}
	}

	for _, board := range r.boards[start:] {
		if board.State == filter.State {
			boards = append(boards, &entity.BoardSummary{
				Id:    board.Id,
				Name:  board.Name,
				State: board.State,
			})
		}
		if len(boards) == pageSize+1 {
			break
		}
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
