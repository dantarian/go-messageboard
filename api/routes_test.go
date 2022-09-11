package api_test

import (
	"pencethren/go-messageboard/api"
	"pencethren/go-messageboard/repositories"
	"testing"

	"github.com/go-chi/chi/v5"
)

// Tests that applying routes to a router doesn't crash and burn.
func TestApplyRoutes(t *testing.T) {
	boardRepo := repositories.NewDefaultBoardRepoMock()
	pingRepo := repositories.NewDefaultPingRepoMock()
	router := api.NewRouter(pingRepo, boardRepo)

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("router.ApplyRoutes panicked unexpectedly: %v", r)
		}
	}()

	router.ApplyRoutes(chi.NewRouter())
}
