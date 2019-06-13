package handler

import (
	"encoding/json"
	"net/http"

	"github.com/wincentrtz/fake-news/domain/postqueue"

	"github.com/gorilla/mux"
)

type PostQueueHandler struct {
	PostQueueUseCase postqueue.Usecase
}

func NewPostQueueHandler(r *mux.Router, us postqueue.Usecase) {
	handler := &PostQueueHandler{
		PostQueueUseCase: us,
	}
	r.HandleFunc("/postqueues", handler.FetchHandler).Methods("GET")
	r.HandleFunc("/postqueues", handler.CreateHandler).Methods("POST")
}

func (pqh *PostQueueHandler) FetchHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := pqh.PostQueueUseCase.FetchPostQueue()
	if err != nil {
		panic("ERROR")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(posts)
}

func (pqh *PostQueueHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := pqh.PostQueueUseCase.CreatePostQueue()
	if err != nil {
		panic("ERROR")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(posts)
}
