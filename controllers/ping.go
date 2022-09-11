package controllers

import (
	"net/http"
	"pencethren/go-messageboard/adapters"
	"pencethren/go-messageboard/operations"
	"pencethren/go-messageboard/repositories"

	"github.com/go-chi/render"
)

type PingPostResponse struct {
	Message string `json:"message"`
}

func (ppr *PingPostResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type PingGetResponse struct {
	Total int `json:"total"`
}

func (ppr *PingGetResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type PingsController struct {
	repository repositories.IPingRepository
}

func NewPingsController(repository repositories.IPingRepository) PingsController {
	return PingsController{repository: repository}
}

func (p *PingsController) PostPing(w http.ResponseWriter, r *http.Request) {
	result := operations.RespondToPing(p.repository)
	render.Status(r, http.StatusOK)
	if err := render.Render(w, r, &PingPostResponse{Message: result}); err != nil {
		adapters.RenderError(w, r, err, http.StatusInternalServerError)
		return
	}
}

func (p *PingsController) GetPings(w http.ResponseWriter, r *http.Request) {
	total := operations.CountPingsReceived(p.repository)

	render.Status(r, http.StatusOK)
	if err := render.Render(w, r, &PingGetResponse{Total: total}); err != nil {
		adapters.RenderError(w, r, err, http.StatusInternalServerError)
		return
	}
}
