package postqueue

import "github.com/wincentrtz/fake-news/models"

type Repository interface {
	FetchPostQueue() ([]*models.PostQueue, error)

	CreatePostQueue() (*models.PostQueue, error)
}
