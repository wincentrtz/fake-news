package usecase

import (
	"time"

	"github.com/wincentrtz/fake-news/domain/post"
	"github.com/wincentrtz/fake-news/models"
)

type postUsecase struct {
	postRepo       post.Repository
	contextTimeout time.Duration
}

func NewPostUsecase(a post.Repository, timeout time.Duration) post.Usecase {
	return &postUsecase{
		postRepo:       a,
		contextTimeout: timeout,
	}
}

func (a *postUsecase) Fetch() ([]*models.Post, error) {
	posts, err := a.postRepo.Fetch()
	if err != nil {
		return nil, err
	}
	return posts, nil
}
