package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"pencethren/go-messageboard/controllers"
	"pencethren/go-messageboard/entities"
	"pencethren/go-messageboard/operations"
	"testing"

	"github.com/gin-gonic/gin"
)

type boardOpsMock struct {
	createBoard func(string, string) (string, error)
}

func (ops *boardOpsMock) CreateBoard(name, description string) (string, error) {
	return ops.createBoard(name, description)
}

func newDefaultBoardOpsMock() *boardOpsMock {
	return &boardOpsMock{
		createBoard: func(name, description string) (string, error) { return "id", nil },
	}
}

func newTestContext() (*httptest.ResponseRecorder, *gin.Context) {
	testRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(testRecorder)

	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	return testRecorder, ctx
}

func mockJsonPost(ctx *gin.Context, content interface{}) {
	ctx.Request.Method = "POST"
	ctx.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
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
			http.StatusOK,
			"{\"id\":\"id\"}",
		},
		"validation failure": {
			map[string]interface{}{"name": "name", "description": "description"},
			func(_, _ string) (string, error) {
				return "", &entities.ValidationError{Msg: "validation error"}
			},
			http.StatusUnprocessableEntity,
			"{\"error\":\"validation error: \"}",
		},
		"business rule failure": {
			map[string]interface{}{"name": "name", "description": "description"},
			func(_, _ string) (string, error) {
				return "", &operations.BusinessRuleError{Msg: "business rule error"}
			},
			http.StatusUnprocessableEntity,
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
			"{\"error\":\"failed to parse request body\"}",
		},
	}

	for scenario, testSpec := range tests {
		t.Run(scenario, func(t *testing.T) {
			boardsOps := newDefaultBoardOpsMock()
			boardsOps.createBoard = testSpec.createBoardOp
			recorder, ctx := newTestContext()

			mockJsonPost(ctx, testSpec.requestBody)

			controller := controllers.NewBoardController(boardsOps)
			controller.PostBoard(ctx)

			body := recorder.Body.String()

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
