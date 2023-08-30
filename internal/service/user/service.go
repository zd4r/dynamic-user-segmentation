package user

import (
	userRepo "github.com/zd4r/dynamic-user-segmentation/internal/repository/user"
)

type Service struct {
	userRepository userRepo.Repository
}

func NewService(userRepository userRepo.Repository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}
