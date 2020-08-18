package server

import (
	"encoding/json"
	"net/http"
)

type handler struct {
	//TODO
	// DB
}

func NewHandler() *handler {
	return &handler{}
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
	// TODO
	respondWith(w, nil, http.StatusCreated)
}
