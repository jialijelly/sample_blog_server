package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jialijelly/sample_blog_server/models"
)

func TestCreateArticleHandler(t *testing.T) {
	testCases := []struct {
		name               string
		arg                models.CreateArticle
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			arg: models.CreateArticle{
				Title:   "Hello World",
				Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				Author:  "John",
			},
			expectedStatusCode: http.StatusCreated,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, err := json.Marshal(tc.arg)
			if err != nil {
				t.Fatalf("Failed to marshal request body : %v", err)
			}
			req, err := http.NewRequest(http.MethodPost, "/articles", bytes.NewReader(body))
			if err != nil {
				t.Fatalf("Failed to create http request : %v", err)
			}
			reqRecorder := httptest.NewRecorder()
			NewHandler().CreateArticle(reqRecorder, req)

			if reqRecorder.Code != tc.expectedStatusCode {
				t.Errorf("Expected %v but got %v", tc.expectedStatusCode, reqRecorder.Code)
			}

		})
	}

	// // Check the response body is what we expect.
	// expected := `{"alive": true}`
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }
}
