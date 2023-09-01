package user

import (
	"context"
	"math/rand"
	"time"

	"github.com/zd4r/dynamic-user-segmentation/internal/errs"
	segmentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/segment"
	userModel "github.com/zd4r/dynamic-user-segmentation/internal/model/user"
)

func (s *Service) GetSegments(ctx context.Context, userId int) ([]segmentModel.Segment, error) {
	return s.userRepository.GetSegments(ctx, userId)
}

func (s *Service) GetPercentOfAllUsers(ctx context.Context, percent int) ([]userModel.User, error) {
	users, err := s.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	rand.New(rand.NewSource(time.Now().Unix()))
	rand.Shuffle(len(users), func(i, j int) {
		users[i], users[j] = users[j], users[i]
	})

	amount := percent * len(users) / 100
	users = users[:amount]

	if len(users) == 0 {
		return nil, errs.ErrNotEnoughUsers
	}

	return users, nil
}
