package postqueue

import (
	"github.com/wincentrtz/fake-news/models"
	"github.com/wincentrtz/fake-news/models/request"
)

type Repository interface {
	FetchPostQueue() ([]*models.PostStatus, error)
	FetchPostQueueById(id int) (*models.PostStatus, error)
	CreatePostQueue(pqreq request.PostQueueRequest) (*models.PostStatus, error)
}
