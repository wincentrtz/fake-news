package usecase

import (
	"time"

	"github.com/wincentrtz/fake-news/models/request"

	"github.com/wincentrtz/fake-news/domain/poststatus"
	"github.com/wincentrtz/fake-news/models"
)

type postStatusUsecase struct {
	postRepo       poststatus.Repository
	contextTimeout time.Duration
}

func NewPostStatusUsecase(pqr poststatus.Repository, timeout time.Duration) poststatus.Usecase {
	return &postStatusUsecase{
		postRepo:       pqr,
		contextTimeout: timeout,
	}
}

func (pqu *postStatusUsecase) FetchPostStatus() ([]*models.PostStatus, error) {
	posts, err := pqu.postRepo.FetchPostStatus()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (pqu *postStatusUsecase) CreatePostStatus(pqreq request.PostStatusRequest) (*models.PostStatus, error) {
	posts, err := pqu.postRepo.CreatePostStatus(pqreq)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (pqu *postStatusUsecase) UpdatePostStatus(id int) (*models.PostStatus, error) {
	posts, err := pqu.postRepo.UpdatePostStatus(id)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
