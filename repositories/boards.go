package repositories

import (
	"pencethren/go-messageboard/entities"

	"github.com/google/uuid"
)

type IBoardRepository interface {
	Add(board *entities.Board) (uuid.UUID, error)
	// Update(board *entities.Board) error
	// Delete(id uuid.UUID) error
	// List(pageSize int, filter *entities.BoardSummary) ([]entities.BoardSummary, error)
	// ListAfter(id uuid.UUID, pageSize int, filter *entities.Board) ([]entities.Board, error)
	// ListBefore(id uuid.UUID, pageSize int, filter *entities.Board) ([]entities.Board, error)
	// Get(id uuid.UUID) (entities.Board, error)
	ExistsWithName(name string) (bool, error)
}
