package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/wincentrtz/fake-news/domain/share"
	"github.com/wincentrtz/fake-news/models"
	"github.com/wincentrtz/fake-news/models/builder"
)

type shareRepository struct {
	Conn *sql.DB
}

func NewShareRepository(Conn *sql.DB) share.Repository {
	return &shareRepository{
		Conn,
	}
}

func (sr *shareRepository) Fetch() ([]*models.Post, error) {
	query := "SELECT id, post_parent_id, post_title, post_description FROM posts"
	rows, err := sr.Conn.Query(query)
	defer rows.Close()

	if err != nil || rows == nil {
		fmt.Println(err)
		return nil, nil
	}

	posts := make([]*models.Post, 0)
	for rows.Next() {
		var id int
		var parent int
		var title string
		var description string
		err = rows.Scan(
			&id,
			&parent,
			&title,
			&description,
		)

		post := builder.NewPost().Id(id).Parent(parent).Title(title).Description(description).Build()

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	return posts, nil
}

func (sr *shareRepository) Create(post *models.Post) (string, error) {
	query := "INSERT INTO posts VALUES (DEFAULT, "+strconv.Itoa(post.Parent)+", '"+post.Title+"', '"+post.Description+"', NOW())"
	rows, err := sr.Conn.Query(query)
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
		return "error", err
	}

	return "success", nil
}
