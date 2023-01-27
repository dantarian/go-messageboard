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
	ListBoards(filter *entity.BoardSearch) ([]*entity.BoardSummary, string, error)
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
	board, err := entity.NewBoard(name, description, entity.BoardStateOpen)
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

func (ops BoardOperations) ListBoards(filter *entity.BoardSearch) ([]*entity.BoardSummary, string, error) {
	logContext := "BoardOperation.ListBoards"
	pageSize := 10

	boards, err := ops.boardsRepo.List(pageSize, filter)
	if err != nil {
		ops.logger.FailedToListBoards(logContext, err)
		return []*entity.BoardSummary{}, "", errors.New("data access error")
	}

	bookmark := ""
	if len(boards) > pageSize {
		bookmark = boards[pageSize].Name
		boards = boards[:pageSize]
	}

	return boards, bookmark, nil
}

func OpenBoard() {}

func CloseBoard() {}

func DeleteBoard() {}
