package report

import (
	"context"
	"time"

	reportModel "github.com/zd4r/dynamic-user-segmentation/internal/model/report"
)

func (s *Service) GetRecordsInIntervalByUser(ctx context.Context, userId int, from, to time.Time) ([]reportModel.Record, error) {
	return s.reportRepository.GetRecordsInIntervalByUser(ctx, userId, from, to)
}
