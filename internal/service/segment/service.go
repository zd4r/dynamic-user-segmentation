package segment

import (
	segmentRepo "github.com/zd4r/dynamic-user-segmentation/internal/repository/segment"
)

type Service struct {
	segmentRepository segmentRepo.Repository
}

func NewService(segmentRepository segmentRepo.Repository) *Service {
	return &Service{
		segmentRepository: segmentRepository,
	}
}
