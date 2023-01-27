package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"pencethren/go-messageboard/controller"
	"pencethren/go-messageboard/entity"
	"pencethren/go-messageboard/operation"
	"strings"
	"testing"
)

type boardOpsMock struct {
	createBoard func(string, string) (string, error)
	listBoards  func(*entity.BoardSearch) ([]*entity.BoardSummary, string, error)
}

func (ops *boardOpsMock) CreateBoard(name, description string) (string, error) {
	return ops.createBoard(name, description)
}

func (ops *boardOpsMock) ListBoards(filter *entity.BoardSearch) ([]*entity.BoardSummary, string, error) {
	return ops.listBoards(filter)
}

func newDefaultBoardOpsMock() *boardOpsMock {
	return &boardOpsMock{
		createBoard: func(name, description string) (string, error) { return "id", nil },
		listBoards: func(filter *entity.BoardSearch) ([]*entity.BoardSummary, string, error) {
			return []*entity.BoardSummary{}, "", nil
		},
	}
}

func newRequest(method string, url string, body map[string]interface{}) (*http.Request, error) {
	var bytes = new(bytes.Buffer)
	if err := json.NewEncoder(bytes).Encode(body); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func TestCreateBoardController(t *testing.T) {
	type testSpec struct {
		requestBody    map[string]interface{}
		createBoardOp  func(string, string) (string, error)
		expectedStatus int
		expectedBody   string
	}
	tests := map[string]testSpec{
		"success": {
			map[string]interface{}{"name": "name", "description": "description"},
			func(_, _ string) (string, error) { return "id", nil },
			http.StatusCreated,
			"{\"id\":\"id\"}",
		},
		"validation failure": {
			map[string]interface{}{"name": "name", "description": "description"},
			func(_, _ string) (string, error) {
				return "", &entity.ValidationError{Msg: "validation error"}
			},
			http.StatusBadRequest,
			"{\"error\":\"validation error: \"}",
		},
		"business rule failure": {
			map[string]interface{}{"name": "name", "description": "description"},
			func(_, _ string) (string, error) {
				return "", &operation.BusinessRuleError{Msg: "business rule error"}
			},
			http.StatusConflict,
			"{\"error\":\"business rule error: \"}",
		},
		"internal server error": {
			map[string]interface{}{"name": "name", "description": "description"},
			func(_, _ string) (string, error) { return "", errors.New("foo") },
			http.StatusInternalServerError,
			"{\"error\":\"foo\"}",
		},
		"request error": {
			map[string]interface{}{"foo": "name", "bar": "description"},
			func(_, _ string) (string, error) { return "", nil },
			http.StatusBadRequest,
			"{\"error\":\"name is required\"}",
		},
	}

	for scenario, testSpec := range tests {
		t.Run(scenario, func(t *testing.T) {
			boardsOps := newDefaultBoardOpsMock()
			boardsOps.createBoard = testSpec.createBoardOp

			req, err := newRequest("POST", "/boards", testSpec.requestBody)
			if err != nil {
				t.Errorf("Failed to create request: %v", err)
			}

			recorder := httptest.NewRecorder()
			controller := controller.NewBoardController(boardsOps)
			handler := http.HandlerFunc(controller.PostBoard)
			handler.ServeHTTP(recorder, req)

			body := strings.Trim(recorder.Body.String(), " \n")

			if recorder.Code != testSpec.expectedStatus || body != testSpec.expectedBody {
				t.Errorf("expected { status: %v, body: %v }, got { status: %v, body: %v }",
					testSpec.expectedStatus,
					testSpec.expectedBody,
					recorder.Code,
					body)
			}
		})
	}
}
