package repository

import (
	"pencethren/go-messageboard/entity"

	"github.com/google/uuid"
)

type IBoardRepository interface {
	Add(board *entity.Board) (uuid.UUID, error)
	// Update(board *entities.Board) error
	// Delete(id uuid.UUID) error
	List(pageSize int, filter *entity.BoardSearch) ([]*entity.BoardSummary, error)
	// Get(id uuid.UUID) (entities.Board, error)
	ExistsWithName(name string) (bool, error)
}
