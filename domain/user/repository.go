package user

import (
	"github.com/wincentrtz/fake-news/models"
)

type Repository interface {
	FetchUserById(userId int) (*models.User, error)
}
