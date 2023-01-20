package repository

import (
	"pencethren/go-messageboard/entity"

	"github.com/google/uuid"
)

type BoardRepoMock struct {
	add            func(*entity.Board) (uuid.UUID, error)
	existsWithName func(string) (bool, error)
	list           func(int, *entity.BoardSearch) ([]*entity.BoardSummary, error)
}

func (br *BoardRepoMock) SetAdd(f func(*entity.Board) (uuid.UUID, error)) {
	br.add = f
}

func (br *BoardRepoMock) SetExistsWithName(f func(string) (bool, error)) {
	br.existsWithName = f
}

func (br *BoardRepoMock) SetList(f func(int, *entity.BoardSearch) ([]*entity.BoardSummary, error)) {
	br.list = f
}

func (br *BoardRepoMock) Add(board *entity.Board) (uuid.UUID, error) {
	return br.add(board)
}

func (br *BoardRepoMock) ExistsWithName(name string) (bool, error) {
	return br.existsWithName(name)
}

func (br *BoardRepoMock) List(pageSize int, boardSearch *entity.BoardSearch) ([]*entity.BoardSummary, error) {
	return br.list(pageSize, boardSearch)
}

func NewDefaultBoardRepoMock() *BoardRepoMock {
	return &BoardRepoMock{
		add:            func(board *entity.Board) (uuid.UUID, error) { return uuid.New(), nil },
		existsWithName: func(name string) (bool, error) { return false, nil },
		list: func(pageSize int, boardSearch *entity.BoardSearch) ([]*entity.BoardSummary, error) {
			return []*entity.BoardSummary{}, nil
		},
	}
}
