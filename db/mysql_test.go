package db

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jialijelly/sample_blog_server/models"
	uuid "github.com/satori/go.uuid"
)

var (
	TestConfig = models.DBConfig{ //TODO: Make test config configurable.
		Type:     "mysql",
		Host:     "localhost",
		Port:     3306,
		Name:     "blog",
		User:     "blog",
		Password: "blog",
	}
	sampleID      = uuid.NewV4().String()
	sampleArticle = models.ArticleInfo{
		ID:      sampleID,
		Title:   "Sample title",
		Content: "Sample content",
		Author:  "Sample author",
	}
)

func testCreateAndGetArticle(t *testing.T, db DB) {
	err := db.CreateArticle(sampleArticle)
	if err != nil {
		t.Fatalf("Failed to create article : %v", err.Response())
	}

	testCases := []struct {
		name           string
		arg            string
		expectedResult models.ArticleInfo
		expectedError  models.Error
	}{
		{
			name:           "Get article successfully",
			arg:            sampleID,
			expectedResult: sampleArticle,
			expectedError:  nil,
		},
		{
			name:          "No article found",
			arg:           "3",
			expectedError: models.NotFound("3"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResult, actualError := db.GetArticle(tc.arg)
			if tc.expectedError != nil {
				if actualError == nil {
					t.Errorf("Expected error but no error returned")
				}
				if actualError.StatusCode() != tc.expectedError.StatusCode() {
					t.Errorf("Expected status %v but got %v", tc.expectedError.StatusCode(), actualError.StatusCode())
				}
				if actualError.Response() != tc.expectedError.Response() {
					t.Errorf("Expected response %+v but got %+v", tc.expectedError.Response(), actualError.Response())
				}
			}
			if actualResult != tc.expectedResult {
				t.Errorf("Expected %+v but got %+v", tc.expectedResult, actualResult)
			}
		})
	}
}

func testListArticles(t *testing.T, db DB) {
	err := db.CreateArticle(sampleArticle)
	if err != nil {
		t.Fatalf("Failed to create article : %v", err)
	}
	sampleID2 := uuid.NewV4().String()
	sampleArticle2 := models.ArticleInfo{
		ID:      sampleID2,
		Title:   "Sample title 2",
		Content: "Sample content 2",
		Author:  "Sample author 2",
	}
	err = db.CreateArticle(sampleArticle2)
	if err != nil {
		t.Fatalf("Failed to create article 2 : %+v", err.Response())
	}
	actual, err := db.ListArticles()
	if err != nil {
		t.Fatalf("Failed to get article : %+v", err.Response())
	}
	expected := []models.ArticleInfo{sampleArticle, sampleArticle2}
	if err := compareArticles(actual, expected); err != nil {
		t.Error(err)
	}
}

func NewTestDB(config models.DBConfig) (DB, error) {
	sqlDB, err := connect(config)
	if err != nil {
		return nil, err
	}
	sqlDB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", config.Name))
	if _, err := sqlDB.Exec(fmt.Sprintf("CREATE DATABASE %s", config.Name)); err != nil {
		return nil, fmt.Errorf("failed to create database : %v", err)
	}
	if _, err := sqlDB.Exec(fmt.Sprintf("USE %s", config.Name)); err != nil {
		return nil, fmt.Errorf("failed to use database : %v", err)
	}
	if err := sqlDB.init(); err != nil {
		return nil, err
	}
	return sqlDB, err
}

func makeArticlesMap(articles []models.ArticleInfo) (map[string]models.ArticleInfo, error) {
	res := make(map[string]models.ArticleInfo)
	for _, article := range articles {
		if _, ok := res[article.ID]; ok {
			return nil, fmt.Errorf("Several articles with the same id: %v", article.ID)
		}
		res[article.ID] = article
	}
	return res, nil
}

func compareArticles(actual, expected []models.ArticleInfo) error {
	actualMap, err := makeArticlesMap(actual)
	if err != nil {
		return err
	}

	expectedMap, err := makeArticlesMap(expected)
	if err != nil {
		return err
	}

	if !reflect.DeepEqual(actualMap, expectedMap) {
		return fmt.Errorf("Articles doesn't match. \n Actual: %v, \n Expected: %v", actual, expected)
	}
	return nil
}

func runTest(t *testing.T, testName string, fn func(*testing.T, DB)) {
	t.Run(testName, func(t *testing.T) {
		sql, err := NewTestDB(TestConfig)
		if err != nil {
			t.Fatalf("Failed to connect to database : %v", err)
		}
		fn(t, sql)
	})
}

func TestBlog(t *testing.T) {
	runTest(t, "testCreateAndGetArticle", testCreateAndGetArticle)
	runTest(t, "testListArticles", testListArticles)
}
