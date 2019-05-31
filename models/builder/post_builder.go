package builder

import "github.com/wincentrtz/fake-news/models"

type postBuilder struct {
	title   string
	author  string
	content string
}

// PostBuilder builder interface
type PostBuilder interface {
	Title(string) PostBuilder
	Author(string) PostBuilder
	Content(string) PostBuilder
	Build() *models.Post
}

// New Builder Initialization
func NewPost() PostBuilder {
	return &postBuilder{}
}

func (ub *postBuilder) Title(title string) PostBuilder {
	ub.title = title
	return ub
}

func (ub *postBuilder) Author(author string) PostBuilder {
	ub.author = author
	return ub
}

func (ub *postBuilder) Content(content string) PostBuilder {
	ub.content = content
	return ub
}

func (ub *postBuilder) Build() *models.Post {
	return &models.Post{
		Title:   ub.title,
		Author:  ub.author,
		Content: ub.content,
	}
}
