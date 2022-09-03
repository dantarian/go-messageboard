package operations_test

import (
	"errors"
	"pencethren/go-messageboard/entities"
	"pencethren/go-messageboard/operations"
	"testing"

	"github.com/google/uuid"
)

type boardRepoMock struct {
	add            func(*entities.Board) (uuid.UUID, error)
	existsWithName func(string) (bool, error)
}

func (br *boardRepoMock) Add(board *entities.Board) (uuid.UUID, error) {
	return br.add(board)
}

func (br *boardRepoMock) ExistsWithName(name string) (bool, error) {
	return br.existsWithName(name)
}

func newDefaultBoardRepoMock() *boardRepoMock {
	return &boardRepoMock{
		add:            func(board *entities.Board) (uuid.UUID, error) { return uuid.New(), nil },
		existsWithName: func(name string) (bool, error) { return false, nil },
	}
}

func TestCreateBoard(t *testing.T) {
	tests := map[string]func(t *testing.T){
		"invalid params":                testCreateBoardInvalidParams,
		"fails to check for name clash": testCreateBoardNameClashCheckError,
		"finds name clash":              testCreateBoardNameClash,
		"fails to persist board":        testCreateBoardFailsToPersist,
		"success":                       testCreateBoardSuccess,
	}

	for scenario, test := range tests {
		t.Run(scenario, func(t *testing.T) {
			test(t)
		})
	}
}

func testCreateBoardInvalidParams(t *testing.T) {
	boardsRepo := newDefaultBoardRepoMock()
	expectedError := "invalid board: name cannot be empty"
	ops := operations.NewBoardOperations(boardsRepo)
	id, err := ops.CreateBoard("", "description")

	if id != "" || err == nil || err.Error() != expectedError {
		t.Errorf("expected (\"\", \"%v\"), got (\"%v\", \"%v\")", expectedError, id, err)
	}
}

func testCreateBoardNameClashCheckError(t *testing.T) {
	boardsRepo := newDefaultBoardRepoMock()
	boardsRepo.existsWithName = func(name string) (bool, error) { return false, errors.New("db error") }
	expectedError := "persistence error"
	ops := operations.NewBoardOperations(boardsRepo)
	id, err := ops.CreateBoard("name", "description")

	if id != "" || err == nil || err.Error() != expectedError {
		t.Errorf("expected (\"\", \"%v\"), got (\"%v\", \"%v\")", expectedError, id, err)
	}
}

func testCreateBoardNameClash(t *testing.T) {
	boardsRepo := newDefaultBoardRepoMock()
	boardsRepo.existsWithName = func(name string) (bool, error) { return true, nil }
	expectedError := "invalid data: a board named 'name' already exists"
	ops := operations.NewBoardOperations(boardsRepo)
	id, err := ops.CreateBoard("name", "description")

	if id != "" || err == nil || err.Error() != expectedError {
		t.Errorf("expected (\"\", \"%v\"), got (\"%v\", \"%v\")", expectedError, id, err)
	}
}

func testCreateBoardFailsToPersist(t *testing.T) {
	boardsRepo := newDefaultBoardRepoMock()
	boardsRepo.add = func(board *entities.Board) (uuid.UUID, error) { return uuid.Nil, errors.New("db error") }
	expectedError := "persistence error"
	ops := operations.NewBoardOperations(boardsRepo)
	id, err := ops.CreateBoard("name", "description")

	if id != "" || err == nil || err.Error() != expectedError {
		t.Errorf("expected (\"\", \"%v\"), got (\"%v\", \"%v\")", expectedError, id, err)
	}
}

func testCreateBoardSuccess(t *testing.T) {
	expectedId := uuid.New()
	boardsRepo := newDefaultBoardRepoMock()
	boardsRepo.add = func(board *entities.Board) (uuid.UUID, error) { return expectedId, nil }
	ops := operations.NewBoardOperations(boardsRepo)
	id, err := ops.CreateBoard("name", "description")

	if id != expectedId.String() || err != nil {
		t.Errorf("expected (\"%v\", nil), got (\"%v\", \"%v\")", expectedId, id, err)
	}
}
