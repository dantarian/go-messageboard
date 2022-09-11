package server

import (
	"fmt"
	"net/http"
	"pencethren/go-messageboard/api"
	"pencethren/go-messageboard/config"
	"pencethren/go-messageboard/data"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	config *config.Config
	router *chi.Mux
}

func NewServer(config *config.Config) *Server {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	pingRepo := data.NewInMemoryPingRepository()
	boardRepo := data.NewInMemoryBoardRepository()

	apiRouter := api.NewRouter(pingRepo, boardRepo)
	apiRouter.ApplyRoutes(router)

	return &Server{config, router}
}

func (s *Server) Run() {
	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.config.Server.Port), s.router); err != nil {
		panic(fmt.Errorf("failed to start server: %w", err))
	}
}
