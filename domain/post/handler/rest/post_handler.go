package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wincentrtz/fake-news/domain/post"
)

type PostHandler struct {
	AUsecase post.Usecase
}

func NewPostHandler(r *mux.Router, us post.Usecase) {
	handler := &PostHandler{
		AUsecase: us,
	}
	r.HandleFunc("/test", TestHandler).Methods("GET")
	r.HandleFunc("/posts", handler.FetchHandler).Methods("GET")
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("asd")
}

// FetchArticle will fetch the article based on given params
func (a *PostHandler) FetchHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := a.AUsecase.Fetch()
	if err != nil {
		panic("ERROR")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(posts)
}
