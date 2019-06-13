package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/wincentrtz/fake-news/domain/post"
	"github.com/wincentrtz/fake-news/models"
	"github.com/wincentrtz/fake-news/models/builder"
	"github.com/wincentrtz/fake-news/models/request"
)

type postRepository struct {
	Conn *sql.DB
}

func NewPostRepository(Conn *sql.DB) post.Repository {
	return &postRepository{
		Conn,
	}
}

func (m *postRepository) FetchPost() ([]*models.Post, error) {
	query := "SELECT * FROM posts"
	rows, err := m.Conn.Query(query)
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
		var date time.Time

		err = rows.Scan(
			&id,
			&parent,
			&title,
			&description,
			&date,
		)

		post := builder.NewPost().
			Id(id).
			Parent(parent).
			Title(title).
			Description(description).
			Date(date).
			Build()

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	return posts, nil
}

func (m *postRepository) CreatePost(pr request.PostRequest) (*models.Post, error) {
	var id int

	query := `INSERT INTO posts
		VALUES (
			DEFAULT,
			` + strconv.Itoa(pr.Parent) + `,
			'` + pr.Title + `',
			'` + pr.Description + `',
			NOW()
		)
		RETURNING id`

	err := m.Conn.QueryRow(query).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	post, err := m.FetchPostById(id)
	if err != nil {
		return nil, err
	}

	return post, err
}

func (m *postRepository) FetchPostById(postId int) (*models.Post, error) {
	var id int
	var parent int
	var title string
	var description string
	var date time.Time

	query := `
		SELECT * FROM posts
		WHERE id =` + strconv.Itoa(postId)
	err := m.Conn.QueryRow(query).Scan(
		&id,
		&parent,
		&title,
		&description,
		&date,
	)

	if err != nil {
		return nil, err
	}

	post := builder.NewPost().
		Id(id).
		Parent(parent).
		Title(title).
		Description(description).
		Date(date).
		Build()

	return post, nil
}
