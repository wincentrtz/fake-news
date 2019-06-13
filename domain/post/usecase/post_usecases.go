package usecase

import (
	"time"

	"github.com/wincentrtz/fake-news/domain/post"
	"github.com/wincentrtz/fake-news/models"
	"github.com/wincentrtz/fake-news/models/request"
)

type postUsecase struct {
	postRepo       post.Repository
	contextTimeout time.Duration
}

func NewPostUsecase(pr post.Repository, timeout time.Duration) post.Usecase {
	return &postUsecase{
		postRepo:       pr,
		contextTimeout: timeout,
	}
}

func (pu *postUsecase) FetchPost() ([]*models.Post, error) {
	posts, err := pu.postRepo.FetchPost()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (pu *postUsecase) CreatePost(pr request.PostRequest) (*models.Post, error) {
	posts, err := pu.postRepo.CreatePost(pr)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
