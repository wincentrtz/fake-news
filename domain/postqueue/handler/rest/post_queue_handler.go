package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wincentrtz/fake-news/domain/postqueue"
	"github.com/wincentrtz/fake-news/models/request"

	"github.com/gorilla/mux"
)

type PostQueueHandler struct {
	PostQueueUseCase postqueue.Usecase
}

func NewPostQueueHandler(r *mux.Router, us postqueue.Usecase) {
	handler := &PostQueueHandler{
		PostQueueUseCase: us,
	}
	r.HandleFunc("/post/queues", handler.FetchHandler).Methods("GET")
	r.HandleFunc("/post/queues", handler.CreateHandler).Methods("POST", "OPTIONS")
}

func (pqh *PostQueueHandler) FetchHandler(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	posts, err := pqh.PostQueueUseCase.FetchPostQueue()
	if err != nil {
		panic("ERROR")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(posts)
}

func (pqh *PostQueueHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	var postQueueRequest request.PostQueueRequest
	_ = json.NewDecoder(r.Body).Decode(&postQueueRequest)
	fmt.Println(postQueueRequest.PostId)
	posts, err := pqh.PostQueueUseCase.CreatePostQueue(postQueueRequest)
	if err != nil {
		panic("ERROR")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
