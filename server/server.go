package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"pencethren/go-messageboard/api"
	"pencethren/go-messageboard/config"
	"pencethren/go-messageboard/data"
	"pencethren/go-messageboard/repository"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	config *config.Config
	router *chi.Mux
}

func NewServer(config *config.Config, db *sql.DB) *Server {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	var boardRepo repository.IBoardRepository
	if db == nil {
		boardRepo = data.NewInMemoryBoardRepository()
	} else {
		boardRepo = data.NewPostgresBoardRepository(db)
	}

	apiRouter := api.NewRouter(boardRepo)
	apiRouter.ApplyRoutes(router)

	return &Server{config, router}
}

func (s *Server) Run() {
	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.config.Server.Port), s.router); err != nil {
		panic(fmt.Errorf("failed to start server: %w", err))
	}
}
