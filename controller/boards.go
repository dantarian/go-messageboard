package controller

import (
	"errors"
	"net/http"
	"pencethren/go-messageboard/adapter"
	"pencethren/go-messageboard/entity"
	"pencethren/go-messageboard/operation"
	"pencethren/go-messageboard/util"
)

type BoardController struct {
	ops    operation.IBoardOperations
	logger *util.ApplicationLogger
}

func NewBoardController(ops operation.IBoardOperations) BoardController {
	return BoardController{ops: ops, logger: util.NewLogger()}
}

func (bc *BoardController) PostBoard(w http.ResponseWriter, r *http.Request) {
	board, err := adapter.NewCreateBoardRequest(r)
	if err != nil {
		bc.logger.Error().Err(err).Msg("failed to parse request body")
		adapter.RenderError(w, r, err, http.StatusBadRequest)
		return
	}

	result, err := bc.ops.CreateBoard(board.Name, board.Description)
	if err != nil {
		var validationError *entity.ValidationError
		var businessRuleError *operation.BusinessRuleError
		switch {
		case errors.As(err, &validationError):
			adapter.RenderError(w, r, err, http.StatusBadRequest)
		case errors.As(err, &businessRuleError):
			adapter.RenderError(w, r, err, http.StatusConflict)
		default:
			adapter.RenderError(w, r, err, http.StatusInternalServerError)
		}
		return
	}

	adapter.RenderCreateBoardResponse(w, r, result)
}

func (bc *BoardController) GetBoards(w http.ResponseWriter, r *http.Request) {
	state, err := entity.ParseBoardState(r.URL.Query().Get("state"))
	if err != nil {
		bc.logger.Error().Err(err).Msg("failed to parse query parameters")
		adapter.RenderError(w, r, err, http.StatusBadRequest)
		return
	}

	searchTerm := r.URL.Query().Get("search")
	bookmark := r.URL.Query().Get("starting_from")

	boardSearch := entity.NewBoardSearch().WithNameContaining(searchTerm).StartingFrom(bookmark)
	if state == entity.BoardStateClosed {
		boardSearch = boardSearch.Closed()
	}

	results, bookmark, err := bc.ops.ListBoards(boardSearch)
	if err != nil {
		adapter.RenderError(w, r, err, http.StatusInternalServerError)
		return
	}

	adapter.RenderBoardsList(w, r, results, bookmark)
}
