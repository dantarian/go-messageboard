package api

import (
	"pencethren/go-messageboard/controllers"
	"pencethren/go-messageboard/operations"
	"pencethren/go-messageboard/repositories"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	boardRepo repositories.IBoardRepository
}

func NewRouter(boardsRepo repositories.IBoardRepository) *Router {
	return &Router{boardsRepo}
}

func (r *Router) ApplyRoutes(chiRouter chi.Router) {
	boardController := controllers.NewBoardController(operations.NewBoardOperations(r.boardRepo))

	chiRouter.Route("/boards", func(r chi.Router) {
		r.Post("/", boardController.PostBoard)
	})
}
