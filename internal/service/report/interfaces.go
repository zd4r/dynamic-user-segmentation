package report

import (
	"context"
	"time"

	reportModel "github.com/zd4r/dynamic-user-segmentation/internal/model/report"
	reportRepo "github.com/zd4r/dynamic-user-segmentation/internal/repository/report"
)

//go:generate mockery --name ReportRepository
type reportRepository interface {
	CreateBatchRecord(ctx context.Context, records []reportModel.Record) error
	GetRecordsInIntervalByUser(ctx context.Context, userId int, from, to time.Time) ([]reportModel.Record, error)
}

var _ reportRepository = (*reportRepo.Repository)(nil)
