package api

import (
	"pencethren/go-messageboard/controllers"
	"pencethren/go-messageboard/operations"
	"pencethren/go-messageboard/repositories"

	"github.com/gin-gonic/gin"
)

type Router struct {
	pingRepo  repositories.IPingRepository
	boardRepo repositories.IBoardRepository
}

func NewRouter(pingRepo repositories.IPingRepository, boardsRepo repositories.IBoardRepository) *Router {
	return &Router{pingRepo, boardsRepo}
}

func (r *Router) ApplyRoutes(engine *gin.Engine) {
	pings := engine.Group("pings")
	pingController := controllers.NewPingsController(r.pingRepo)
	pingController.ApplyRoutes(pings)

	boards := engine.Group("boards")
	boardsController := controllers.NewBoardController(operations.NewBoardOperations(r.boardRepo))
	boardsController.ApplyRoutes(boards)
}
