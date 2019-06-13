package usecase

import (
	"time"

	"github.com/wincentrtz/fake-news/models/request"

	"github.com/wincentrtz/fake-news/domain/postqueue"
	"github.com/wincentrtz/fake-news/models"
)

type postQueueUsecase struct {
	postRepo       postqueue.Repository
	contextTimeout time.Duration
}

func NewPostQueueUsecase(pqr postqueue.Repository, timeout time.Duration) postqueue.Usecase {
	return &postQueueUsecase{
		postRepo:       pqr,
		contextTimeout: timeout,
	}
}

func (pqu *postQueueUsecase) FetchPostQueue() ([]*models.PostQueue, error) {
	posts, err := pqu.postRepo.FetchPostQueue()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (pqu *postQueueUsecase) CreatePostQueue(pqreq request.PostQueueRequest) (*models.PostQueue, error) {
	posts, err := pqu.postRepo.CreatePostQueue(pqreq)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
