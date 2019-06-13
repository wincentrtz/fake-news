package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/wincentrtz/fake-news/models/request"

	"github.com/wincentrtz/fake-news/domain/poststatus"
	"github.com/wincentrtz/fake-news/models"
	"github.com/wincentrtz/fake-news/models/builder"
)

type postStatusRepository struct {
	Conn *sql.DB
}

func NewPostStatusRepository(Conn *sql.DB) poststatus.Repository {
	return &postStatusRepository{Conn}
}

func (m *postStatusRepository) FetchPostStatus() ([]*models.PostStatus, error) {
	query := "SELECT post_status.id, post_id,post_title, status FROM post_status JOIN posts ON (post_status.post_id = posts.id)"
	rows, err := m.Conn.Query(query)
	defer rows.Close()
	if err != nil || rows == nil {
		fmt.Println(err)
		return nil, nil
	}
	posts := make([]*models.PostStatus, 0)

	for rows.Next() {
		var id int
		var postID int
		var postTitle string
		var status int
		err = rows.Scan(
			&id,
			&postID,
			&postTitle,
			&status,
		)

		if err != nil {
			return nil, err
		}

		post := builder.NewPostStatus().Id(id).PostId(postID).PostTitle(postTitle).Status(status).Build()

		posts = append(posts, post)
	}

	return posts, nil
}

func (m *postStatusRepository) CreatePostStatus(pqreq request.PostStatusRequest) (*models.PostStatus, error) {

	var id int

	fmt.Println(pqreq.PostId)
	query := `INSERT INTO post_status (post_id, status, created_on)
		VALUES($1,$2,$3)
		RETURNING id
	`
	err := m.Conn.QueryRow(query, &pqreq.PostId, 0, time.Now()).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	postStatus, err := m.FetchPostStatusById(id)

	if err != nil {
		return nil, err
	}

	return postStatus, nil
}

func (m *postStatusRepository) FetchPostStatusById(id int) (*models.PostStatus, error) {

	var postID int
	var postTitle string
	var status int

	query := `
		SELECT post_id,post_title, status FROM 
		post_status JOIN posts ON (post_status.post_id = posts.id)
		where post_status.id =` + strconv.Itoa(id)
	err := m.Conn.QueryRow(query).Scan(&postID, &postTitle, &status)

	if err != nil {
		return nil, err
	}

	post := builder.NewPostStatus().Id(id).PostId(postID).PostTitle(postTitle).Status(status).Build()

	return post, nil
}

func (m *postStatusRepository) UpdatePostStatus(id int) (*models.PostStatus, error) {
	postStatus, err := m.FetchPostStatusById(id)
	if err != nil {
		return nil, err
	}

	if postStatus.Status != 2 {
		postStatus.Status = postStatus.Status + 1

		query := `UPDATE post_status
			SET status = $1
			WHERE id = $2`

		_, err = m.Conn.Exec(query, postStatus.Status, postStatus.Id)

		if err != nil {
			return nil, err
		}
	}

	return postStatus, nil
}
