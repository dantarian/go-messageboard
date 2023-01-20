package operation_test

import (
	"errors"
	"pencethren/go-messageboard/entity"
	"pencethren/go-messageboard/operation"
	"pencethren/go-messageboard/repository"
	"testing"

	"github.com/google/uuid"
)

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
	boardsRepo := repository.NewDefaultBoardRepoMock()
	expectedError := "invalid board: name cannot be empty"
	ops := operation.NewBoardOperations(boardsRepo)
	id, err := ops.CreateBoard("", "description")

	if id != "" || err == nil || err.Error() != expectedError {
		t.Errorf("expected (\"\", \"%v\"), got (\"%v\", \"%v\")", expectedError, id, err)
	}
}

func testCreateBoardNameClashCheckError(t *testing.T) {
	boardsRepo := repository.NewDefaultBoardRepoMock()
	boardsRepo.SetExistsWithName(func(name string) (bool, error) { return false, errors.New("db error") })
	expectedError := "persistence error"
	ops := operation.NewBoardOperations(boardsRepo)
	id, err := ops.CreateBoard("name", "description")

	if id != "" || err == nil || err.Error() != expectedError {
		t.Errorf("expected (\"\", \"%v\"), got (\"%v\", \"%v\")", expectedError, id, err)
	}
}

func testCreateBoardNameClash(t *testing.T) {
	boardsRepo := repository.NewDefaultBoardRepoMock()
	boardsRepo.SetExistsWithName(func(name string) (bool, error) { return true, nil })
	expectedError := "invalid data: a board named 'name' already exists"
	ops := operation.NewBoardOperations(boardsRepo)
	id, err := ops.CreateBoard("name", "description")

	if id != "" || err == nil || err.Error() != expectedError {
		t.Errorf("expected (\"\", \"%v\"), got (\"%v\", \"%v\")", expectedError, id, err)
	}
}

func testCreateBoardFailsToPersist(t *testing.T) {
	boardsRepo := repository.NewDefaultBoardRepoMock()
	boardsRepo.SetAdd(func(board *entity.Board) (uuid.UUID, error) { return uuid.Nil, errors.New("db error") })
	expectedError := "persistence error"
	ops := operation.NewBoardOperations(boardsRepo)
	id, err := ops.CreateBoard("name", "description")

	if id != "" || err == nil || err.Error() != expectedError {
		t.Errorf("expected (\"\", \"%v\"), got (\"%v\", \"%v\")", expectedError, id, err)
	}
}

func testCreateBoardSuccess(t *testing.T) {
	expectedId := uuid.New()
	boardsRepo := repository.NewDefaultBoardRepoMock()
	boardsRepo.SetAdd(func(board *entity.Board) (uuid.UUID, error) { return expectedId, nil })
	ops := operation.NewBoardOperations(boardsRepo)
	id, err := ops.CreateBoard("name", "description")

	if id != expectedId.String() || err != nil {
		t.Errorf("expected (\"%v\", nil), got (\"%v\", \"%v\")", expectedId, id, err)
	}
}
