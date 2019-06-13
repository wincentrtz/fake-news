package usecase

import (
	"time"

	"github.com/wincentrtz/fake-news/domain/user"
	"github.com/wincentrtz/fake-news/models"
)

type userUsecase struct {
	userRepo       user.Repository
	contextTimeout time.Duration
}

func NewUserUsecase(ur user.Repository, timeout time.Duration) user.Usecase {
	return &userUsecase{
		userRepo:       ur,
		contextTimeout: timeout,
	}
}

func (pu *userUsecase) FetchUserById(userId int) (*models.User, error) {
	user, err := pu.userRepo.FetchUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
