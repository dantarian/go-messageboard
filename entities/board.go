package entities

import (
	"github.com/google/uuid"
)

type BoardState int64

const (
	BoardClosed BoardState = iota
	BoardOpen
)

type BoardSummary struct {
	Id    uuid.UUID
	Name  string
	State BoardState
}

type Board struct {
	Id          uuid.UUID
	Name        string
	Description string
	State       BoardState
}

func NewBoard(name string, description string, state BoardState) (*Board, error) {
	board := Board{Id: uuid.New(), Name: name, Description: description, State: state}

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
