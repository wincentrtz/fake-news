package share

import (
	"github.com/wincentrtz/fake-news/models"
)

type Repository interface {
	Fetch() ([]*models.Post, error)
	Create(*models.Post) (string, error)
}
