package builder

import "github.com/wincentrtz/fake-news/models"

type postBuilder struct {
	id          int
	parent      int
	title       string
	description string
}

// PostBuilder builder interface
type PostBuilder interface {
	Id(int) PostBuilder
	Parent(int) PostBuilder
	Title(string) PostBuilder
	Description(string) PostBuilder
	Build() *models.Post
}

// New Builder Initialization
func NewPost() PostBuilder {
	return &postBuilder{}
}

func (ub *postBuilder) Id(id int) PostBuilder {
	ub.id = id
	return ub
}

func (ub *postBuilder) Parent(parent int) PostBuilder {
	ub.parent = parent
	return ub
}

func (ub *postBuilder) Title(title string) PostBuilder {
	ub.title = title
	return ub
}

func (ub *postBuilder) Description(description string) PostBuilder {
	ub.description = description
	return ub
}

func (ub *postBuilder) Build() *models.Post {
	return &models.Post{
		Id:          ub.id,
		Parent:      ub.parent,
		Title:       ub.title,
		Description: ub.description,
	}
}
