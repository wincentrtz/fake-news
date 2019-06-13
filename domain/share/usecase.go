package share

import (
	"github.com/wincentrtz/fake-news/models"
)

type Usecase interface {
	Fetch() ([]*models.Post, error)
	Create(post *models.Post) (string, error)
}
