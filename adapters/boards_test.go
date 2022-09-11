package adapters_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"pencethren/go-messageboard/adapters"
	"testing"
)

func TestCreateBoardRequest(t *testing.T) {
	type test struct {
		body           map[string]interface{}
		expectedResult *adapters.CreateBoardRequest
		expectedError  string
	}

	tests := map[string]test{
		"name and description binds": {map[string]interface{}{"name": "name", "description": "description"}, &adapters.CreateBoardRequest{"name", "description"}, ""},
		"missing description binds":  {map[string]interface{}{"name": "name"}, &adapters.CreateBoardRequest{"name", ""}, ""},
		"missing name errors":        {map[string]interface{}{"description": "description"}, nil, "name is required"},
	}

	for scenario, details := range tests {
		t.Run(scenario, func(t *testing.T) {
			var bytes = new(bytes.Buffer)
			if err := json.NewEncoder(bytes).Encode(details.body); err != nil {
				t.Errorf("Failed to encode request body: %v", err)
				return
			}

			req, err := http.NewRequest("POST", "/boards", bytes)
			if err != nil {
				t.Errorf("Failed to create request: %v", err)
				return
			}
			req.Header.Set("Content-Type", "application/json")

			result, err := adapters.NewCreateBoardRequest(req)

			if (result == nil && details.expectedResult != nil) ||
				(result != nil && details.expectedResult == nil) ||
				(result != nil && details.expectedResult != nil && *result != *details.expectedResult) ||
				(err != nil && err.Error() != details.expectedError) ||
				(err == nil && details.expectedError != "") {
				t.Errorf("expected (%p, %v); got (%p, %v)", details.expectedResult, details.expectedError, result, err)
			}
		})
	}
}
