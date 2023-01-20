package operation

import (
	"errors"
	"fmt"
	"pencethren/go-messageboard/entity"
	"pencethren/go-messageboard/repository"
	"pencethren/go-messageboard/util"
)

type IBoardOperations interface {
	CreateBoard(name, description string) (string, error)
}

type BoardOperations struct {
	boardsRepo repository.IBoardRepository
	logger     *util.ApplicationLogger
}

func NewBoardOperations(boardsRepo repository.IBoardRepository) BoardOperations {
	return BoardOperations{boardsRepo: boardsRepo, logger: util.NewLogger()}
}

func (ops BoardOperations) CreateBoard(name, description string) (string, error) {
	logContext := "BoardOperations.CreateBoard"

	// Validate
	board, err := entity.NewBoard(name, description, entity.BoardOpen)
	if err != nil {
		ops.logger.FailedToInstantiateBoard(logContext, err)
		return "", err
	}

	// Business rule - boards must have unique names
	nameAlreadyUsed, err := ops.boardsRepo.ExistsWithName(name)
	if err != nil {
		ops.logger.FailedToCheckForBoardNameReuse(logContext, err)
		return "", errors.New("persistence error")
	}
	if nameAlreadyUsed {
		ops.logger.DuplicateBoardName(logContext, name)
		return "", NewBusinessRuleError("invalid data", fmt.Sprintf("a board named '%v' already exists", name))
	}

	// Persist board
	id, err := ops.boardsRepo.Add(board)
	if err != nil {
		ops.logger.FailedToPersistBoard(logContext, err)
		return "", errors.New("persistence error")
	}

	return fmt.Sprint(id), nil
}

func (ops BoardOperations) ListOpenBoards(filter *entity.BoardSearch) ([]*entity.BoardSummary, string, error) {
	logContext := "BoardOperation.ListOpenBoards"
	pageSize := 10

	boards, err := ops.boardsRepo.List(pageSize, filter)
	if err != nil {
		ops.logger.FailedToListBoards(logContext, err)
		return []*entity.BoardSummary{}, "", errors.New("data access error")
	}

	bookmark := ""
	if len(boards) > pageSize {
		bookmark = boards[pageSize-1].Name
	}

	if filter.Order == entity.Descending {
		size := pageSize
		if len(boards) < pageSize {
			size = len(boards)
		}
		boards = reverse(boards[:size])
	}

	return boards, bookmark, nil
}

func reverse[T any](original []T) (reversed []T) {
	reversed = make([]T, len(original))
	copy(reversed, original)

	for i := len(reversed)/2 - 1; i >= 0; i-- {
		tmp := len(reversed) - 1 - i
		reversed[i], reversed[tmp] = reversed[tmp], reversed[i]
	}

	return
}

func ListClosedBoards() {}

func OpenBoard() {}

func CloseBoard() {}

func DeleteBoard() {}
