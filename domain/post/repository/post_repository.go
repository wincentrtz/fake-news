package repository

import (
	"database/sql"
	"fmt"

	"github.com/wincentrtz/fake-news/domain/post"
	"github.com/wincentrtz/fake-news/models"
	"github.com/wincentrtz/fake-news/models/builder"
)

type postRepository struct {
	Conn *sql.DB
}

func NewPostRepository(Conn *sql.DB) post.Repository {
	return &postRepository{
		Conn,
	}
}

func (m *postRepository) Fetch() ([]*models.Post, error) {
	query := "SELECT id, post_title, post_description FROM posts"
	rows, err := m.Conn.Query(query)
	defer rows.Close()
	if err != nil || rows == nil {
		fmt.Println(err)
		return nil, nil
	}
	posts := make([]*models.Post, 0)
	for rows.Next() {
		var id int
		var title string
		var description string
		err = rows.Scan(
			&id,
			&title,
			&description,
		)

		post := builder.NewPost().Id(id).Title(title).Description(description).Build()

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	return posts, nil
}
