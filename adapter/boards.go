package adapter

import (
	"errors"
	"net/http"
	"pencethren/go-messageboard/entity"

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

type BoardSummary struct {
	Id    string `json:"id" binding:"required"`
	Name  string `json:"name"`
	State string `json:"state"`
}

type ListBoardsResponse struct {
	Boards []BoardSummary `json:"boards"`
	Next   string         `json:"next"`
}

func (res *ListBoardsResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	return nil
}

func RenderBoardsList(w http.ResponseWriter, r *http.Request, list []*entity.BoardSummary, bookmark string) {
	summaries := []BoardSummary{}
	for _, summary := range list {
		summaries = append(summaries, BoardSummary{
			Id:    summary.Id.String(),
			Name:  summary.Name,
			State: summary.State.String(),
		})
	}

	url := ""
	if bookmark != "" {
		newUrl := *r.URL
		query := newUrl.Query()
		query.Set("starting_from", bookmark)
		newUrl.RawQuery = query.Encode()
		url = newUrl.String()
	}

	_ = render.Render(w, r, &ListBoardsResponse{summaries, url})
}
