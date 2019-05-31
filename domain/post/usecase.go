package post

import (
	"github.com/wincentrtz/fake-news/models"
)

// Usecase represent the article's usecases
type Usecase interface {
	Fetch() ([]*models.Post, error)
}
