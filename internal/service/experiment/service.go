package experiment

import (
	experimentRepo "github.com/zd4r/dynamic-user-segmentation/internal/repository/experiment"
)

type Service struct {
	experimentRepository *experimentRepo.Repository
}

func NewService(experimentRepository *experimentRepo.Repository) *Service {
	return &Service{
		experimentRepository: experimentRepository,
	}
}
