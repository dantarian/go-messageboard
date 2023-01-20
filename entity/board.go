package entity

import (
	"github.com/google/uuid"
)

type BoardState int64

const (
	BoardClosed BoardState = iota
	BoardOpen
)

type SortOrder int64

const (
	Ascending SortOrder = iota
	Descending
)

type BoardSummary struct {
	Id    uuid.UUID
	Name  string
	State BoardState
}

type Board struct {
	BoardSummary
	Description string
}

type BoardSearch struct {
	SearchTerm string
	State      BoardState
	Order      SortOrder
	Bookmark   string
}

func NewBoard(name string, description string, state BoardState) (*Board, error) {
	board := Board{
		BoardSummary: BoardSummary{
			Id:    uuid.New(),
			Name:  name,
			State: state,
		},
		Description: description,
	}

	if err := board.Validate(); err != nil {
		return nil, err
	}

	return &board, nil
}

func (b *Board) Validate() error {
	if b.Name == "" {
		return &ValidationError{Msg: "invalid board", reasons: []string{"name cannot be empty"}}
	}

	return nil
}

func NewBoardSearch() *BoardSearch {
	return &BoardSearch{State: BoardOpen, Order: Ascending}
}

func (boardSearch *BoardSearch) WithNameContaining(term string) *BoardSearch {
	boardSearch.SearchTerm = term
	return boardSearch
}

func (boardSearch *BoardSearch) Closed() *BoardSearch {
	boardSearch.State = BoardClosed
	return boardSearch
}

func (boardSearch *BoardSearch) Descending() *BoardSearch {
	boardSearch.Order = Descending
	return boardSearch
}

func (boardSearch *BoardSearch) StartingFrom(name string) *BoardSearch {
	boardSearch.Bookmark = name
	return boardSearch
}
