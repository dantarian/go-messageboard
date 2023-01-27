package entity

import (
	"fmt"

	"github.com/google/uuid"
)

type BoardState int64

const (
	BoardStateClosed BoardState = iota
	BoardStateOpen
)

func (bs BoardState) String() string {
	return []string{"closed", "open"}[bs]
}

func ParseBoardState(s string) (BoardState, error) {
	switch s {
	case "closed":
		return BoardStateClosed, nil
	case "open":
		return BoardStateOpen, nil
	default:
		return BoardStateClosed, fmt.Errorf("unrecognised board state: %s", s)
	}
}

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
	return &BoardSearch{State: BoardStateOpen}
}

func (boardSearch *BoardSearch) WithNameContaining(term string) *BoardSearch {
	boardSearch.SearchTerm = term
	return boardSearch
}

func (boardSearch *BoardSearch) Closed() *BoardSearch {
	boardSearch.State = BoardStateClosed
	return boardSearch
}

func (boardSearch *BoardSearch) StartingFrom(name string) *BoardSearch {
	boardSearch.Bookmark = name
	return boardSearch
}
