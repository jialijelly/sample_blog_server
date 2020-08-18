package db

import (
	"github.com/jialijelly/sample_blog_server/models"
)

// Using an interface to support different database type in the future.
type DB interface {
	CreateArticle(article models.ArticleInfo) models.Error
	ListArticles() ([]models.ArticleInfo, models.Error)
	GetArticle(id string) (models.ArticleInfo, models.Error)
}
