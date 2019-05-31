package builder

import "github.com/wincentrtz/fake-news/models"

type authorBuilder struct {
	name string
}

// AuthorBuilder builder interface
type AuthorBuilder interface {
	Name(string) AuthorBuilder
	Build() *models.User
}

// New Builder Initialization
func NewAuthor() AuthorBuilder {
	return &authorBuilder{}
}

func (ab *authorBuilder) Name(name string) AuthorBuilder {
	ab.name = name
	return ab
}

func (ab *authorBuilder) Build() *models.User {
	return &models.User{
		Name: ab.name,
	}
}
