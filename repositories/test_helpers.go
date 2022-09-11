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

type PingRepoMock struct {
	add   func(entities.Ping)
	count func() int
}

func (pr *PingRepoMock) SetAdd(f func(entities.Ping)) {
	pr.add = f
}

func (pr *PingRepoMock) SetCount(f func() int) {
	pr.count = f
}

func (pr *PingRepoMock) Add(p entities.Ping) {
	pr.add(p)
}

func (pr *PingRepoMock) Count() int {
	return pr.count()
}

func NewDefaultPingRepoMock() *PingRepoMock {
	return &PingRepoMock{
		add:   func(_ entities.Ping) {},
		count: func() int { return 0 },
	}
}
