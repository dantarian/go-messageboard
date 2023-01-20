package entity_test

import (
	"pencethren/go-messageboard/entity"
	"testing"

	"github.com/google/uuid"
)

func TestNewBoard(t *testing.T) {
	type test struct {
		name               string
		description        string
		state              entity.BoardState
		expectBoard        bool
		expectErrorMessage string
	}

	tests := map[string]test{
		"full open board validates":                 {"some name", "some description", entity.BoardOpen, true, ""},
		"open board with no name does not validate": {"", "some description", entity.BoardOpen, false, "invalid board: name cannot be empty"},
		"open board with no description validates":  {"some name", "", entity.BoardOpen, true, ""},
		"full closed board validates":               {"some name", "some description", entity.BoardClosed, true, ""},
	}

	for scenario, details := range tests {
		t.Run(scenario, func(t *testing.T) {
			board, err := entity.NewBoard(details.name, details.description, details.state)

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
	board := entity.Board{
		BoardSummary: entity.BoardSummary{
			Id:    uuid.New(),
			Name:  "Some name",
			State: entity.BoardOpen},
		Description: "Some description.",
	}

	if err := board.Validate(); err != nil {
		t.Errorf("Expected board to validate; got: %v", err)
	}
}

func TestBoardValidateBoardWithoutNameIsInvalid(t *testing.T) {
	board := entity.Board{
		BoardSummary: entity.BoardSummary{
			Id:    uuid.New(),
			Name:  "",
			State: entity.BoardOpen,
		},
		Description: "Some description.",
	}

	err := board.Validate()
	if err == nil {
		t.Error("Expected board to be invalid with no name, but validation passed.")
	}

}

func TestBoardSearch(t *testing.T) {
	type test struct {
		search   *entity.BoardSearch
		expected *entity.BoardSearch
	}

	tests := map[string]test{
		"base search":      {entity.NewBoardSearch(), &entity.BoardSearch{"", entity.BoardOpen, entity.Ascending, ""}},
		"with search term": {entity.NewBoardSearch().WithNameContaining("foo"), &entity.BoardSearch{"foo", entity.BoardOpen, entity.Ascending, ""}},
		"closed boards":    {entity.NewBoardSearch().Closed(), &entity.BoardSearch{"", entity.BoardClosed, entity.Ascending, ""}},
		"descending order": {entity.NewBoardSearch().Descending(), &entity.BoardSearch{"", entity.BoardOpen, entity.Descending, ""}},
		"from name":        {entity.NewBoardSearch().StartingFrom("foo"), &entity.BoardSearch{"", entity.BoardOpen, entity.Ascending, "foo"}},
		"chaining":         {entity.NewBoardSearch().WithNameContaining("foo").Closed().Descending().StartingFrom("bar"), &entity.BoardSearch{"foo", entity.BoardClosed, entity.Descending, "bar"}},
	}

	for scenario, details := range tests {
		t.Run(scenario, func(t *testing.T) {
			if *details.search != *details.expected {
				t.Errorf("%v: expected %v, got %v", scenario, details.expected, details.search)
			}
		})
	}
}
