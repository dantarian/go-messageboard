package server

import (
	"fmt"
	"net/http"
	"pencethren/go-messageboard/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config *config.Config
	router *gin.Engine
}

func NewServer(config *config.Config) *Server {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return &Server{config, router}
}

func RunServer(s *Server) {
	if err := s.router.Run(fmt.Sprintf(":%d", s.config.Server.Port)); err != nil {
		panic(fmt.Errorf("failed to start server: %w", err))
	}
}
