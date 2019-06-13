package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wincentrtz/fake-news/domain/share"
	"github.com/wincentrtz/fake-news/models"
)

type ShareHandler struct {
	ShareUsecase share.Usecase
}

func NewShareHandler(r *mux.Router, us share.Usecase) {
	handler := &ShareHandler{
		ShareUsecase: us,
	}

	r.HandleFunc("/api/share", handler.FetchHandler).Methods("GET")
	r.HandleFunc("/api/share", handler.CreateHandler).Methods("POST")
}

func (sh *ShareHandler) FetchHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := sh.ShareUsecase.Fetch()
	if err != nil {
		panic("ERROR")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(posts)
}

func (sh *ShareHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	post := &models.Post{}
	err = json.Unmarshal(b, post)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	result, err := sh.ShareUsecase.Create(post)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"code": `+strconv.Itoa(http.StatusInternalServerError)+`, "description": `+err.Error()+`}`)
	} else {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"code": `+strconv.Itoa(http.StatusOK)+`, "description": `+result+`}`)
	}
}
