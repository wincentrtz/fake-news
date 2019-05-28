package builder

import "github.com/wincentrtz/fake-news/models"

type userBuilder struct {
	name string
}

// UserBuilder builder interface
type UserBuilder interface {
	Name(string) UserBuilder
	Build() *models.User
}

// New Builder Initialization
func NewUser() UserBuilder {
	return &userBuilder{}
}

func (ub *userBuilder) Name(name string) UserBuilder {
	ub.name = name
	return ub
}

func (ub *userBuilder) Build() *models.User {
	return &models.User{
		Name: ub.name,
	}
}
