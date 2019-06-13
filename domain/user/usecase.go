package user

import (
	"github.com/wincentrtz/fake-news/models"
)

type Usecase interface {
	FetchUserById(userId int) (*models.User, error)
}
