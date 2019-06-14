package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/wincentrtz/fake-news/domain/chart"
	"github.com/wincentrtz/fake-news/models"
	"github.com/wincentrtz/fake-news/models/builder"
)

type chartRepository struct {
	Conn *sql.DB
}

func NewChartRepository(Conn *sql.DB) chart.Repository {
	return &chartRepository{
		Conn,
	}
}

func (m *chartRepository) FetchChart() ([]*models.Chart, error) {
	query := `
		SELECT 
			DATE(posts.created_on), 
			post_status.status, 
			COUNT(posts.id) 
		FROM 
			posts
		FULL JOIN 
			post_status ON posts.post_parent_id = post_status.post_id
		GROUP BY 
			DATE(posts.created_on), 
			post_status.status 
		ORDER BY 
			DATE(posts.created_on),
			post_status.status`

	rows, err := m.Conn.Query(query)
	defer rows.Close()
	if err != nil || rows == nil {
		fmt.Println(err)
		return nil, nil
	}

	charts := make([]*models.Chart, 0)
	curr := models.Chart{}

	if rows.Next() {
		var firstDate time.Time
		var firstCount int
		var firstStatus sql.NullInt64

		err = rows.Scan(
			&firstDate,
			&firstStatus,
			&firstCount,
		)
		if err != nil {
			return nil, err
		}

		if !firstStatus.Valid {
			curr.Date = firstDate
			curr.Status.Unclassified = firstCount
		} else if firstStatus.Int64 == 0 {
			curr.Date = firstDate
			curr.Status.Hoax = firstCount
		} else if firstStatus.Int64 == 1 {
			curr.Date = firstDate
			curr.Status.Valid = firstCount
		}
	}

	for rows.Next() {
		var date time.Time
		var count int
		var status sql.NullInt64

		err = rows.Scan(
			&date,
			&status,
			&count,
		)
		if err != nil {
			return nil, err
		}

		if date == curr.Date {
			if !status.Valid {
				curr.Status.Unclassified = count
			} else if status.Int64 == 0 {
				curr.Status.Hoax = count
			} else if status.Int64 == 1 {
				curr.Status.Valid = count
			}
		} else {
			chartStatus := builder.NewChartStatus().
				Hoax(curr.Status.Hoax).
				Valid(curr.Status.Valid).
				Unclassified(curr.Status.Unclassified).
				Build()

			chart := builder.NewChart().
				Date(curr.Date).
				Status(*chartStatus).
				Build()

			charts = append(charts, chart)

			curr = models.Chart{}
			if !status.Valid {
				curr.Date = date
				curr.Status.Unclassified = count
			} else if status.Int64 == 0 {
				curr.Date = date
				curr.Status.Hoax = count
			} else if status.Int64 == 1 {
				curr.Date = date
				curr.Status.Valid = count
			}
		}
	}

	{
		chartStatus := builder.NewChartStatus().
			Hoax(curr.Status.Hoax).
			Valid(curr.Status.Valid).
			Unclassified(curr.Status.Unclassified).
			Build()

		chart := builder.NewChart().
			Date(curr.Date).
			Status(*chartStatus).
			Build()

		charts = append(charts, chart)
	}

	return charts, nil
}
