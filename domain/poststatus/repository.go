package poststatus

import (
	"github.com/wincentrtz/fake-news/models"
	"github.com/wincentrtz/fake-news/models/request"
)

type Repository interface {
	FetchPostStatus() ([]*models.PostStatus, error)
	FetchPostStatusById(id int) (*models.PostStatus, error)
	CreatePostStatus(pqreq request.PostStatusRequest) (*models.PostStatus, error)
	UpdatePostStatus(id int) (*models.PostStatus, error)
}
