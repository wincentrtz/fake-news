package repository

import (
	"database/sql"
	"fmt"

	"github.com/wincentrtz/fake-news/domain/postqueue"
	"github.com/wincentrtz/fake-news/models"
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
		t := new(models.PostQueue)
		err = rows.Scan(
			&t.Id,
			&t.PostId,
			&t.PostTitle,
			&t.Progress,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, t)
	}

	println(posts)
	return posts, nil
}
