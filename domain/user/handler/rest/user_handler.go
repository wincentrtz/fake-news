package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/wincentrtz/fake-news/domain/user"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserUsecase user.Usecase
}

func NewUserHandler(r *mux.Router, us user.Usecase) {
	handler := &UserHandler{
		UserUsecase: us,
	}
	r.HandleFunc("/api/user/{id}", handler.FetchByIdHandler).Methods("GET")
}

func (uh *UserHandler) FetchByIdHandler(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	vars := mux.Vars(r)
	i, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic("ERROR")
	}
	user, err := uh.UserUsecase.FetchUserById(i)
	if err != nil {
		panic("ERROR")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
