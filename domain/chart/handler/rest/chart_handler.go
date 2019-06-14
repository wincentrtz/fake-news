package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wincentrtz/fake-news/domain/chart"
)

type ChartHandler struct {
	ChartUsecase chart.Usecase
}

func NewChartHandler(r *mux.Router, cu chart.Usecase) {
	handler := &ChartHandler{
		ChartUsecase: cu,
	}

	r.HandleFunc("/api/chart", handler.FetchHandler).Methods("GET")
}

func (ch *ChartHandler) FetchHandler(w http.ResponseWriter, r *http.Request) {
	charts, err := ch.ChartUsecase.FetchChart()
	if err != nil {
		panic("ERROR")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(charts)
}
