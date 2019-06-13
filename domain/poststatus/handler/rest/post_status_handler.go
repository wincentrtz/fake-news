package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wincentrtz/fake-news/domain/poststatus"
	"github.com/wincentrtz/fake-news/models/request"

	"github.com/gorilla/mux"
)

type PostStatusHandler struct {
	PostStatusUseCase poststatus.Usecase
}

func NewPostStatusHandler(r *mux.Router, us poststatus.Usecase) {
	handler := &PostStatusHandler{
		PostStatusUseCase: us,
	}
	r.HandleFunc("/post/status", handler.FetchHandler).Methods("GET")
	r.HandleFunc("/post/status", handler.CreateHandler).Methods("POST", "OPTIONS")
}

func (pqh *PostStatusHandler) FetchHandler(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	posts, err := pqh.PostStatusUseCase.FetchPostStatus()
	if err != nil {
		panic("ERROR")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(posts)
}

func (pqh *PostStatusHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	var postStatusRequest request.PostStatusRequest
	_ = json.NewDecoder(r.Body).Decode(&postStatusRequest)
	fmt.Println(postStatusRequest.PostId)
	posts, err := pqh.PostStatusUseCase.CreatePostStatus(postStatusRequest)
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
