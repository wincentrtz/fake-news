package postqueue

import (
	"github.com/wincentrtz/fake-news/models"
)

// Usecase represent the article's usecases
type Usecase interface {
	FetchPostQueue() ([]*models.PostQueue, error)
}
