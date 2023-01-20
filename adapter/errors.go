package adapter

import (
	"net/http"

	"github.com/go-chi/render"
)

type errorResponse struct {
	Err            error  `json:"-"`
	HttpStatusCode int    `json:"-"`
	ErrorText      string `json:"error,omitempty"`
}

func (e *errorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HttpStatusCode)
	return nil
}

func RenderError(w http.ResponseWriter, r *http.Request, err error, statusCode int) {
	_ = render.Render(w, r, &errorResponse{err, statusCode, err.Error()})
}
