package post

import (
	"github.com/wincentrtz/fake-news/models"
)

type Usecase interface {
	Fetch() ([]*models.Post, error)
}
