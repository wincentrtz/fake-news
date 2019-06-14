package chart

import (
	"github.com/wincentrtz/fake-news/models"
)

type Usecase interface {
	FetchChart() ([]*models.Chart, error)
}
