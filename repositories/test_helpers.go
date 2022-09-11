package repositories

import (
	"pencethren/go-messageboard/entities"

	"github.com/google/uuid"
)

type BoardRepoMock struct {
	add            func(*entities.Board) (uuid.UUID, error)
	existsWithName func(string) (bool, error)
}

func (br *BoardRepoMock) SetAdd(f func(*entities.Board) (uuid.UUID, error)) {
	br.add = f
}

func (br *BoardRepoMock) SetExistsWithName(f func(string) (bool, error)) {
	br.existsWithName = f
}

func (br *BoardRepoMock) Add(board *entities.Board) (uuid.UUID, error) {
	return br.add(board)
}

func (br *BoardRepoMock) ExistsWithName(name string) (bool, error) {
	return br.existsWithName(name)
}

func NewDefaultBoardRepoMock() *BoardRepoMock {
	return &BoardRepoMock{
		add:            func(board *entities.Board) (uuid.UUID, error) { return uuid.New(), nil },
		existsWithName: func(name string) (bool, error) { return false, nil },
	}
}
