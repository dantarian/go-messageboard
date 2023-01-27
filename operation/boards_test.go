package operation_test

import (
	"errors"
	"fmt"
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

func TestListBoards(t *testing.T) {
	tests := map[string]func(t *testing.T){
		"fails to retrieve boards":     testListBoardsFailsToRead,
		"no boards available":          testListBoardsEmptyList,
		"fewer than _pageSize_ boards": testListBoardsSmallList,
		"more than _pageSize_ boards":  testListBoardsLargeList,
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

func testListBoardsFailsToRead(t *testing.T) {
	boardsRepo := repository.NewDefaultBoardRepoMock()
	boardsRepo.SetList(func(i int, bs *entity.BoardSearch) ([]*entity.BoardSummary, error) {
		return []*entity.BoardSummary{}, errors.New("db error")
	})
	expectedError := "data access error"
	ops := operation.NewBoardOperations(boardsRepo)
	list, bookmark, err := ops.ListBoards(&entity.BoardSearch{})

	if len(list) != 0 || bookmark != "" || err.Error() != expectedError {
		t.Errorf("expected ([], \"\", \"%v\"), got (%v, \"%v\", \"%v\")", expectedError, list, bookmark, err)
	}
}

func testListBoardsEmptyList(t *testing.T) {
	boardsRepo := repository.NewDefaultBoardRepoMock()
	boardsRepo.SetList(func(i int, bs *entity.BoardSearch) ([]*entity.BoardSummary, error) {
		return []*entity.BoardSummary{}, nil
	})
	ops := operation.NewBoardOperations(boardsRepo)
	list, bookmark, err := ops.ListBoards(&entity.BoardSearch{})

	if len(list) != 0 || bookmark != "" || err != nil {
		t.Errorf("expected ([], \"\", nil), got (%v, \"%v\", \"%v\")", list, bookmark, err)
	}
}

func testListBoardsSmallList(t *testing.T) {
	boardsRepo := repository.NewDefaultBoardRepoMock()
	boardsRepo.SetList(func(i int, bs *entity.BoardSearch) ([]*entity.BoardSummary, error) {
		return []*entity.BoardSummary{
			{
				Id:    uuid.New(),
				Name:  "Test board",
				State: entity.BoardStateOpen,
			},
		}, nil
	})
	ops := operation.NewBoardOperations(boardsRepo)
	list, bookmark, err := ops.ListBoards(&entity.BoardSearch{})

	if len(list) != 1 || bookmark != "" || err != nil {
		t.Errorf("expected ([<Test board>], \"\", nil), got (%v, \"%v\", \"%v\")", list, bookmark, err)
	}
}

func testListBoardsLargeList(t *testing.T) {
	boardsRepo := repository.NewDefaultBoardRepoMock()
	boardsRepo.SetList(func(i int, bs *entity.BoardSearch) ([]*entity.BoardSummary, error) {
		boards := []*entity.BoardSummary{}

		for i := 0; i < 11; i++ {
			boards = append(boards, &entity.BoardSummary{
				Id:    uuid.New(),
				Name:  fmt.Sprintf("Test board %v", i),
				State: entity.BoardStateOpen,
			})
		}

		return boards, nil
	})
	ops := operation.NewBoardOperations(boardsRepo)
	list, bookmark, err := ops.ListBoards(&entity.BoardSearch{})

	if len(list) != 10 || bookmark != "Test board 11" || err != nil {
		t.Errorf("expected ([<Test board 1>, ..., <Test board 10>], \"Test board 11\", nil), got (%v, \"%v\", \"%v\")", list, bookmark, err)
	}
}
