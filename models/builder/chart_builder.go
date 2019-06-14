package builder

import (
	"time"

	"github.com/wincentrtz/fake-news/models"
)

type chartBuilder struct {
	date   time.Time
	status models.ChartStatus
}

type chartStatusBuilder struct {
	valid        int
	hoax         int
	unclassified int
}

type ChartBuilder interface {
	Date(time.Time) ChartBuilder
	Status(models.ChartStatus) ChartBuilder
	Build() *models.Chart
}

type ChartStatusBuilder interface {
	Valid(int) ChartStatusBuilder
	Hoax(int) ChartStatusBuilder
	Unclassified(int) ChartStatusBuilder
	Build() *models.ChartStatus
}

func NewChart() ChartBuilder {
	return &chartBuilder{}
}

func NewChartStatus() ChartStatusBuilder {
	return &chartStatusBuilder{}
}

func (cb *chartBuilder) Date(date time.Time) ChartBuilder {
	cb.date = date
	return cb
}

func (cb *chartBuilder) Status(status models.ChartStatus) ChartBuilder {
	cb.status = status
	return cb
}

func (cb *chartBuilder) Build() *models.Chart {
	return &models.Chart{
		Date:   cb.date,
		Status: cb.status,
	}
}

func (csb *chartStatusBuilder) Valid(valid int) ChartStatusBuilder {
	csb.valid = valid
	return csb
}

func (csb *chartStatusBuilder) Hoax(hoax int) ChartStatusBuilder {
	csb.hoax = hoax
	return csb
}

func (csb *chartStatusBuilder) Unclassified(unclassified int) ChartStatusBuilder {
	csb.unclassified = unclassified
	return csb
}

func (csb *chartStatusBuilder) Build() *models.ChartStatus {
	return &models.ChartStatus{
		Valid:        csb.valid,
		Hoax:         csb.hoax,
		Unclassified: csb.unclassified,
	}
}
