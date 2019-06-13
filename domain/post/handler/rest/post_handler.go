package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wincentrtz/fake-news/domain/post"
	"github.com/wincentrtz/fake-news/models/request"
)

type PostHandler struct {
	PostUsecase post.Usecase
}

func NewPostHandler(r *mux.Router, us post.Usecase) {
	handler := &PostHandler{
		PostUsecase: us,
	}

	r.HandleFunc("/api/post", handler.FetchHandler).Methods("GET")
	r.HandleFunc("/api/post", handler.CreateHandler).Methods("POST")
}

func (ph *PostHandler) FetchHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := ph.PostUsecase.FetchPost()
	if err != nil {
		panic("ERROR")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(posts)
}

func (ph *PostHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	var postRequest request.PostRequest
	_ = json.NewDecoder(r.Body).Decode(&postRequest)

	posts, err := ph.PostUsecase.CreatePost(postRequest)
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
