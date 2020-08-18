package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jialijelly/sample_blog_server/db"
	"github.com/jialijelly/sample_blog_server/models"
)

const (
	Success = "Success"
)

type handler struct {
	DB db.DB
}

func NewHandler(database db.DB) *handler {
	return &handler{DB: database}
}

func respondWith(w http.ResponseWriter, data interface{}, statusCode int) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(dataJSON)
}

// CreateArticle will validate user input article and store in database.
func (h *handler) CreateArticle(w http.ResponseWriter, req *http.Request) {
	article := models.CreateArticle{}
	err := json.NewDecoder(req.Body).Decode(&article)
	defer req.Body.Close()
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		errType := models.FailedToReadRequest()
		respondWith(w, errType.Response(), errType.StatusCode())
		return
	}
	if errType := article.Validate(); errType != nil {
		respondWith(w, errType.Response(), errType.StatusCode())
		return
	}

	newArticle := article.Register()
	if errType := h.DB.CreateArticle(newArticle); errType != nil {
		respondWith(w, errType.Response(), errType.StatusCode())
		return
	}

	respondWith(w, models.Response{
		Status:  http.StatusCreated,
		Message: Success,
		Data:    models.CreateArticleResponse{ID: newArticle.ID},
	}, http.StatusCreated)
}

// ListArticles will list all the articles found in the database.
func (h *handler) ListArticles(w http.ResponseWriter, req *http.Request) {
	articles, err := h.DB.ListArticles()
	if err != nil {
		respondWith(w, err.Response(), err.StatusCode())
		return
	}

	respondWith(w, models.Response{
		Status:  http.StatusOK,
		Message: Success,
		Data:    articles,
	}, http.StatusOK)
}

// GetArticle will retrieve the article based on the id defined in the request.
func (h *handler) GetArticle(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["article_id"]

	article, err := h.DB.GetArticle(id)
	if err != nil {
		respondWith(w, err.Response(), err.StatusCode())
		return
	}

	respondWith(w, models.Response{
		Status:  http.StatusOK,
		Message: Success,
		Data:    article,
	}, http.StatusOK)
}
