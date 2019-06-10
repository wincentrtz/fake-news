package repository

import (
	"database/sql"
	"fmt"

	"github.com/wincentrtz/fake-news/domain/postqueue"
	"github.com/wincentrtz/fake-news/models"
	"github.com/wincentrtz/fake-news/models/builder"
)

type postQueueRepository struct {
	Conn *sql.DB
}

func NewPostQueueRepository(Conn *sql.DB) postqueue.Repository {
	return &postQueueRepository{
		Conn,
	}
}

func (m *postQueueRepository) FetchPostQueue() ([]*models.PostQueue, error) {
	query := "SELECT post_queues.post_id, post_id,post_title, progress FROM post_queues JOIN posts ON (post_queues.post_id = posts.id)"
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

		post := builder.NewPostQueue().Id(id).PostId(postID).PostTitle(postTitle).Progress(progress).Build()

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	println(posts)
	return posts, nil
}
