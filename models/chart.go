package models

import (
	"time"
)

type Chart struct {
	Date   time.Time   `json:"date"`
	Status ChartStatus `json:"status"`
}

type ChartStatus struct {
	Valid        int `json:"valid"`
	Hoax         int `json:"hoax"`
	Unclassified int `json:"unclassified"`
}
