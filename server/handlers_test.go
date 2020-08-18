package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/jialijelly/sample_blog_server/models"
)

var (
	sampleArticle = models.ArticleInfo{
		ID:      "1",
		Title:   "Sample title 1",
		Content: "Sample content 1",
		Author:  "Sample author 1",
	}

	sampleListArticles = []models.ArticleInfo{
		sampleArticle,
		{
			ID:      "2",
			Title:   "Sample title 2",
			Content: "Sample content 2",
			Author:  "Sample author 2",
		},
	}
)

type mockDB struct{}

func (m *mockDB) CreateArticle(article models.ArticleInfo) models.Error {
	return nil
}

func (m *mockDB) ListArticles() ([]models.ArticleInfo, models.Error) {
	return sampleListArticles, nil
}

func (m *mockDB) GetArticle(id string) (models.ArticleInfo, models.Error) {
	return sampleArticle, nil
}

func TestCreateArticleHandler(t *testing.T) {
	testCases := []struct {
		name               string
		arg                models.CreateArticle
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "test with missing title",
			arg: models.CreateArticle{
				Content: "Sample content",
				Author:  "Sample author",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"status":400,"message":"Missing required field 'title' in request body.","data":null}`,
		},
		{
			name: "test with missing content",
			arg: models.CreateArticle{
				Title:  "Sample title",
				Author: "Sample author",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"status":400,"message":"Missing required field 'content' in request body.","data":null}`,
		},
		{
			name: "test with missing author",
			arg: models.CreateArticle{
				Title:   "Sample title",
				Content: "Sample content",
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"status":400,"message":"Missing required field 'author' in request body.","data":null}`,
		},
		{
			name: "test with success",
			arg: models.CreateArticle{
				Title:   "Sample title",
				Content: "Sample content",
				Author:  "Sample author",
			},
			expectedStatusCode: http.StatusCreated,
			expectedResponse:   `{"status":201,"message":"Success","data":{"id":".*"}}`,
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
			NewHandler(&mockDB{}).CreateArticle(reqRecorder, req)

			if reqRecorder.Code != tc.expectedStatusCode {
				t.Errorf("Expected %v but got %v", tc.expectedStatusCode, reqRecorder.Code)
			}

			matched, err := regexp.MatchString(tc.expectedResponse, reqRecorder.Body.String())
			if err != nil {
				t.Fatalf("Failed to validate response : %v", err)
			}
			if !matched {
				t.Errorf("Expected like %v but got %v", tc.expectedResponse, reqRecorder.Body.String())
			}
		})
	}
}

func TestListArticles(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/articles", nil)
	if err != nil {
		t.Fatalf("Failed to create http request : %v", err)
	}

	reqRecorder := httptest.NewRecorder()
	NewHandler(&mockDB{}).ListArticles(reqRecorder, req)

	expectedStatusCode := http.StatusOK
	expectedResponse := models.Response{
		Status:  expectedStatusCode,
		Message: Success,
		Data:    sampleListArticles,
	}
	if reqRecorder.Code != expectedStatusCode {
		t.Errorf("Expected %v but got %v", expectedStatusCode, reqRecorder.Code)
	}
	expected, err := json.Marshal(expectedResponse)
	if err != nil {
		t.Fatalf("Failed to marshal expected response : %v", err)
	}
	if reqRecorder.Body.String() != string(expected) {
		t.Errorf("Expected %v but got %v", expected, reqRecorder.Body.String())
	}
}

func TestGetArticle(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/articles/1", nil)
	if err != nil {
		t.Fatalf("Failed to create http request : %v", err)
	}

	reqRecorder := httptest.NewRecorder()
	NewHandler(&mockDB{}).GetArticle(reqRecorder, req)

	expectedStatusCode := http.StatusOK
	expectedResponse := models.Response{
		Status:  expectedStatusCode,
		Message: Success,
		Data:    sampleArticle,
	}
	if reqRecorder.Code != expectedStatusCode {
		t.Errorf("Expected %v but got %v", expectedStatusCode, reqRecorder.Code)
	}
	expected, err := json.Marshal(expectedResponse)
	if err != nil {
		t.Fatalf("Failed to marshal expected response : %v", err)
	}
	if reqRecorder.Body.String() != string(expected) {
		t.Errorf("Expected %v but got %v", expected, reqRecorder.Body.String())
	}
}
