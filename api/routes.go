package api

import (
	"pencethren/go-messageboard/controller"
	"pencethren/go-messageboard/operation"
	"pencethren/go-messageboard/repository"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	boardRepo repository.IBoardRepository
}

func NewRouter(boardsRepo repository.IBoardRepository) *Router {
	return &Router{boardsRepo}
}

func (r *Router) ApplyRoutes(chiRouter chi.Router) {
	boardController := controller.NewBoardController(operation.NewBoardOperations(r.boardRepo))

	chiRouter.Route("/boards", func(r chi.Router) {
		r.Post("/", boardController.PostBoard)
	})
}
