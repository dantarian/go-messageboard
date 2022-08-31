package server

import (
	"fmt"
	"pencethren/go-messageboard/api"
	"pencethren/go-messageboard/config"
	"pencethren/go-messageboard/data"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config *config.Config
	router *gin.Engine
}

func NewServer(config *config.Config) *Server {
	router := gin.Default()

	pingRepo := data.NewInMemoryPingRepository()
	boardRepo := data.NewInMemoryBoardRepository()

	apiRouter := api.NewRouter(pingRepo, boardRepo)
	apiRouter.ApplyRoutes(router)

	return &Server{config, router}
}

func (s *Server) Run() {
	if err := s.router.Run(fmt.Sprintf(":%d", s.config.Server.Port)); err != nil {
		panic(fmt.Errorf("failed to start server: %w", err))
	}
}
