package segment

import (
	"context"

	"github.com/zd4r/dynamic-user-segmentation/internal/errs"
)

func (s *Service) DeleteBySlug(ctx context.Context, slug string) error {
	rowsAffected, err := s.segmentRepository.DeleteBySlug(ctx, slug)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errs.ErrSegmentNotFound
	}

	return nil
}
