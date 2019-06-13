package usecase

import (
	"time"

	"github.com/wincentrtz/fake-news/domain/share"
	"github.com/wincentrtz/fake-news/models"
)

type shareUsecase struct {
	shareRepo      share.Repository
	contextTimeout time.Duration
}

func NewShareUsecase(repository share.Repository, timeout time.Duration) share.Usecase {
	return &shareUsecase{
		shareRepo:      repository,
		contextTimeout: timeout,
	}
}

func (su *shareUsecase) Fetch() ([]*models.Post, error) {
	posts, err := su.shareRepo.Fetch()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (su *shareUsecase) Create(post *models.Post) (string, error) {
	result, err := su.shareRepo.Create(post)
	return result, err
}
