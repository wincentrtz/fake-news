package builder

import "github.com/wincentrtz/fake-news/models"

type userBuilder struct {
	id       int
	name     string
	email    string
	password string
}

// UserBuilder builder interface
type UserBuilder interface {
	Id(int) UserBuilder
	Name(string) UserBuilder
	Email(string) UserBuilder
	Password(string) UserBuilder
	Build() *models.User
}

// New Builder Initialization
func NewUser() UserBuilder {
	return &userBuilder{}
}

func (ub *userBuilder) Id(id int) UserBuilder {
	ub.id = id
	return ub
}

func (ub *userBuilder) Email(email string) UserBuilder {
	ub.email = email
	return ub
}

func (ub *userBuilder) Name(name string) UserBuilder {
	ub.name = name
	return ub
}

func (ub *userBuilder) Password(password string) UserBuilder {
	ub.password = password
	return ub
}

func (ub *userBuilder) Build() *models.User {
	return &models.User{
		Id:       ub.id,
		Email:    ub.email,
		Password: ub.password,
		Name:     ub.name,
	}
}
