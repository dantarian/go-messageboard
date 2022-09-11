package controllers

import (
	"errors"
	"net/http"
	"pencethren/go-messageboard/adapters"
	"pencethren/go-messageboard/entities"
	"pencethren/go-messageboard/operations"
	"pencethren/go-messageboard/util"
)

type BoardController struct {
	ops    operations.IBoardOperations
	logger *util.ApplicationLogger
}

func NewBoardController(ops operations.IBoardOperations) BoardController {
	return BoardController{ops: ops, logger: util.NewLogger()}
}

func (bc *BoardController) PostBoard(w http.ResponseWriter, r *http.Request) {
	board, err := adapters.NewCreateBoardRequest(r)
	if err != nil {
		bc.logger.Error().Err(err).Msg("failed to parse request body")
		adapters.RenderError(w, r, err, http.StatusBadRequest)
		return
	}

	result, err := bc.ops.CreateBoard(board.Name, board.Description)
	if err != nil {
		var validationError *entities.ValidationError
		var businessRuleError *operations.BusinessRuleError
		switch {
		case errors.As(err, &validationError):
			adapters.RenderError(w, r, err, http.StatusBadRequest)
		case errors.As(err, &businessRuleError):
			adapters.RenderError(w, r, err, http.StatusConflict)
		default:
			adapters.RenderError(w, r, err, http.StatusInternalServerError)
		}
		return
	}

	adapters.RenderCreateBoardResponse(w, r, result)
}
