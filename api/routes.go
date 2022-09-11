package api

import (
	"pencethren/go-messageboard/controllers"
	"pencethren/go-messageboard/operations"
	"pencethren/go-messageboard/repositories"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	pingRepo  repositories.IPingRepository
	boardRepo repositories.IBoardRepository
}

func NewRouter(pingRepo repositories.IPingRepository, boardsRepo repositories.IBoardRepository) *Router {
	return &Router{pingRepo, boardsRepo}
}

func (r *Router) ApplyRoutes(chiRouter chi.Router) {
	pingsController := controllers.NewPingsController(r.pingRepo)
	boardController := controllers.NewBoardController(operations.NewBoardOperations(r.boardRepo))

	chiRouter.Route("/pings", func(r chi.Router) {
		r.Post("/", pingsController.PostPing)
		r.Get("/", pingsController.GetPings)
	})

	chiRouter.Route("/boards", func(r chi.Router) {
		r.Post("/", boardController.PostBoard)
	})
}
