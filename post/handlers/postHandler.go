package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewPostHandler(r *mux.Router) {
	r.HandleFunc("/test", TestHandler).Methods("GET")
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("asd")
}
