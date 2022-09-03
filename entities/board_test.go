package entities_test

import (
	"pencethren/go-messageboard/entities"
	"testing"

	"github.com/google/uuid"
)

func TestNewBoard(t *testing.T) {
	type test struct {
		name               string
		description        string
		state              entities.BoardState
		expectBoard        bool
		expectErrorMessage string
	}

	tests := map[string]test{
		"full open board validates":                 {"some name", "some description", entities.BoardOpen, true, ""},
		"open board with no name does not validate": {"", "some description", entities.BoardOpen, false, "invalid board: name cannot be empty"},
		"open board with no description validates":  {"some name", "", entities.BoardOpen, true, ""},
		"full closed board validates":               {"some name", "some description", entities.BoardClosed, true, ""},
	}

	for scenario, details := range tests {
		t.Run(scenario, func(t *testing.T) {
			board, err := entities.NewBoard(details.name, details.description, details.state)

			if (board != nil) != details.expectBoard {
				var expectation string
				if details.expectBoard {
					expectation = "want board"
				} else {
					expectation = "want no board"
				}
				t.Errorf("%v: %v, got %v", scenario, expectation, board)
			}

			if details.expectErrorMessage == "" {
				if err != nil {
					t.Errorf("%v: expected no error, got %v", scenario, err)
				}
			} else {
				if err == nil {
					t.Errorf("%v: expected error, got nil", scenario)
				} else if details.expectErrorMessage != err.Error() {
					t.Errorf("%v: expected error '%v', got '%v'", scenario, details.expectErrorMessage, err)
				}
			}
		})
	}

}

func TestBoardValidateFullBoardIsValid(t *testing.T) {
	board := entities.Board{Id: uuid.New(), Name: "Some name", Description: "Some description.", State: entities.BoardOpen}

	if err := board.Validate(); err != nil {
		t.Errorf("Expected board to validate; got: %v", err)
	}
}

func TestBoardValidateBoardWithoutNameIsInvalid(t *testing.T) {
	board := entities.Board{Id: uuid.New(), Name: "", Description: "Some description.", State: entities.BoardOpen}

	err := board.Validate()
	if err == nil {
		t.Error("Expected board to be invalid with no name, but validation passed.")
	}

}
