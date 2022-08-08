package server

import (
	"fmt"
	"pencethren/go-messageboard/api"
	"pencethren/go-messageboard/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config *config.Config
	router *gin.Engine
}

func NewServer(config *config.Config, buildRoutes api.RouteBuilder) *Server {
	router := gin.Default()

	buildRoutes(router)

	return &Server{config, router}
}

func (s *Server) Run() {
	if err := s.router.Run(fmt.Sprintf(":%d", s.config.Server.Port)); err != nil {
		panic(fmt.Errorf("failed to start server: %w", err))
	}
}
