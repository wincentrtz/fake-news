package usecase

import (
	"time"

	"github.com/wincentrtz/fake-news/domain/postqueue"
	"github.com/wincentrtz/fake-news/models"
)

type postQueueUsecase struct {
	postRepo       postqueue.Repository
	contextTimeout time.Duration
}

func NewPostQueueUsecase(a postqueue.Repository, timeout time.Duration) postqueue.Usecase {
	return &postQueueUsecase{
		postRepo:       a,
		contextTimeout: timeout,
	}
}

func (a *postQueueUsecase) FetchPostQueue() ([]*models.PostQueue, error) {
	posts, err := a.postRepo.FetchPostQueue()
	if err != nil {
		return nil, err
	}
	return posts, nil
}
