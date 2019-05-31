package repository

import (
	"github.com/wincentrtz/fake-news/domain/post"
	"github.com/wincentrtz/fake-news/models"
)

type postRepository struct{}

func NewPostRepository() post.Repository {
	return &postRepository{}
}

func (m *postRepository) Fetch() ([]*models.Post, error) {
	posts := make([]*models.Post, 0)
	for i := 0; i < 5; i++ {
		posts = append(posts, &models.Post{})
	}
	return posts, nil
}
