package report

import (
	"context"

	reportModel "github.com/zd4r/dynamic-user-segmentation/internal/model/report"
)

func (s *Service) CreateBatchRecord(ctx context.Context, records []reportModel.Record) error {
	if len(records) != 0 {
		return s.reportRepository.CreateBatchRecord(ctx, records)
	}

	return nil
}
