package models

// ArticleInfo provide details of an article.
type ArticleInfo struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// CreateArticle is the structure of request body required for CreateArticle request.
type CreateArticle struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
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
