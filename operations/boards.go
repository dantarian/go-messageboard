package operations

import (
	"errors"
	"fmt"
	"pencethren/go-messageboard/entities"
	"pencethren/go-messageboard/repositories"
	"pencethren/go-messageboard/util"
)

type BoardOperations struct {
	boardsRepo repositories.IBoardRepository
	logger     *util.ApplicationLogger
}

func NewBoardOperations(boardsRepo repositories.IBoardRepository) BoardOperations {
	return BoardOperations{boardsRepo: boardsRepo, logger: util.NewLogger()}
}

func (ops BoardOperations) CreateBoard(name, description string) (string, error) {
	logContext := "BoardOperations.CreateBoard"

	// Validate
	board, err := entities.NewBoard(name, description, entities.BoardOpen)
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

func ListOpenBoards() {}

func ListClosedBoards() {}

func OpenBoard() {}

func CloseBoard() {}

func DeleteBoard() {}
