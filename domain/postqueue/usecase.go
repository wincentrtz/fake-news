package postqueue

import (
	"github.com/wincentrtz/fake-news/models"
	"github.com/wincentrtz/fake-news/models/request"
)

// Usecase represent the article's usecases
type Usecase interface {
	FetchPostQueue() ([]*models.PostStatus, error)

	CreatePostQueue(pqreq request.PostQueueRequest) (*models.PostStatus, error)
}
