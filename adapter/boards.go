package adapter

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type CreateBoardRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func (board *CreateBoardRequest) Bind(r *http.Request) error {
	if board.Name == "" {
		return errors.New("name is required")
	}

	return nil
}

func NewCreateBoardRequest(r *http.Request) (*CreateBoardRequest, error) {
	var createBoardRequest CreateBoardRequest
	if err := render.Bind(r, &createBoardRequest); err != nil {
		return nil, err
	}

	return &createBoardRequest, nil
}

type CreateBoardResponse struct {
	Id string `json:"id"`
}

func (res *CreateBoardResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusCreated)
	return nil
}

func RenderCreateBoardResponse(w http.ResponseWriter, r *http.Request, id string) {
	_ = render.Render(w, r, &CreateBoardResponse{Id: id})
}
