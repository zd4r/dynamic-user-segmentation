package segment

import (
	"context"
)

func (s *Service) Create(ctx context.Context, slug string) error {
	return s.segmentRepository.Create(ctx, slug)
}
