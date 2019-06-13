package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/wincentrtz/fake-news/models/request"

	"github.com/wincentrtz/fake-news/domain/postqueue"
	"github.com/wincentrtz/fake-news/models"
	"github.com/wincentrtz/fake-news/models/builder"
)

type postQueueRepository struct {
	Conn *sql.DB
}

func NewPostQueueRepository(Conn *sql.DB) postqueue.Repository {
	return &postQueueRepository{Conn}
}

func (m *postQueueRepository) FetchPostQueue() ([]*models.PostQueue, error) {
	query := "SELECT post_queues.id, post_id,post_title, progress FROM post_queues JOIN posts ON (post_queues.post_id = posts.id)"
	rows, err := m.Conn.Query(query)
	defer rows.Close()
	if err != nil || rows == nil {
		fmt.Println(err)
		return nil, nil
	}
	posts := make([]*models.PostQueue, 0)

	for rows.Next() {
		var id int
		var postID int
		var postTitle string
		var progress int
		err = rows.Scan(
			&id,
			&postID,
			&postTitle,
			&progress,
		)

		if err != nil {
			return nil, err
		}

		post := builder.NewPostQueue().Id(id).PostId(postID).PostTitle(postTitle).Progress(progress).Build()

		posts = append(posts, post)
	}

	return posts, nil
}

func (m *postQueueRepository) CreatePostQueue(pqreq request.PostQueueRequest) (*models.PostQueue, error) {

	var id int

	fmt.Println(pqreq.PostId)
	query := `INSERT INTO post_queues (post_id, progress, created_on)
		VALUES($1,$2,$3)
		RETURNING id
	`
	err := m.Conn.QueryRow(query, &pqreq.PostId, 0, time.Now()).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	postQueue := builder.NewPostQueue().Id(int(id)).Build()

	return postQueue, nil
}
