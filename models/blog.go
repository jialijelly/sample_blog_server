package models

import (
	uuid "github.com/satori/go.uuid"
)

// ArticleInfo provide details of an article.
type ArticleInfo struct {
	ID      string `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
	Author  string `json:"author" db:"author"`
}

// CreateArticle is the structure of request body required for CreateArticle request.
type CreateArticle struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (a CreateArticle) Validate() Error {
	if a.Title == "" {
		return MissingRequiredParam("title")
	}
	if a.Content == "" {
		return MissingRequiredParam("content")
	}
	if a.Author == "" {
		return MissingRequiredParam("author")
	}
	return nil
}

func (a CreateArticle) Register() ArticleInfo {
	return ArticleInfo{
		ID: uuid.NewV4().String(),
		Title: a.Title,
		Content: a.Content,
		Author: a.Author,
	}
}

// CreateArticleResponse is the result of CreateArticle request.
type CreateArticleResponse struct {
	ID string `json:"id"`
}

// Response is a generic API response.
type Response struct {
	Status  int         `json:"status"`  // HTTP status code.
	Message string      `json:"message"` // Error description or status description.
	Data    interface{} `json:"data"`    // Request result, if applicable.
}
