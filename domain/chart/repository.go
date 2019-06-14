package chart

import (
	"github.com/wincentrtz/fake-news/models"
)

type Repository interface {
	FetchChart() ([]*models.Chart, error)
}
