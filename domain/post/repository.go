package post

import (
	"github.com/wincentrtz/fake-news/models"
)

type Repository interface {
	Fetch() ([]*models.Post, error)
}
