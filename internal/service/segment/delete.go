package segment

import (
	"context"

	"github.com/zd4r/dynamic-user-segmentation/internal/errs"
	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"
)

func (s *Service) Delete(ctx context.Context, segment *segmentModel.Segment) error {
	rowsAffected, err := s.segmentRepository.Delete(ctx, segment)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}
