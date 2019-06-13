package poststatus

import (
	"github.com/wincentrtz/fake-news/models"
	"github.com/wincentrtz/fake-news/models/request"
)

// Usecase represent the article's usecases
type Usecase interface {
	FetchPostStatus() ([]*models.PostStatus, error)
	CreatePostStatus(pqreq request.PostStatusRequest) (*models.PostStatus, error)
	UpdatePostStatus(id int) (*models.PostStatus, error)
}
