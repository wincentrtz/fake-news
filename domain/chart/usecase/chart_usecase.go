package usecase

import (
	"time"

	"github.com/wincentrtz/fake-news/domain/chart"
	"github.com/wincentrtz/fake-news/models"
)

type chartUsecase struct {
	chartRepo      chart.Repository
	contextTimeout time.Duration
}

func NewChartUsecase(cr chart.Repository, timeout time.Duration) chart.Usecase {
	return &chartUsecase{
		chartRepo:      cr,
		contextTimeout: timeout,
	}
}

func (cu *chartUsecase) FetchChart() ([]*models.Chart, error) {
	charts, err := cu.chartRepo.FetchChart()
	if err != nil {
		return nil, err
	}
	return charts, nil
}
