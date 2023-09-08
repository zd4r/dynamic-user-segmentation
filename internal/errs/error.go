package errs

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrSegmentNotFound     = errors.New("segment not found")
	ErrExperimentNotFound  = errors.New("experiment not found")
	ErrInvalidUserId       = errors.New("invalid user id")
	ErrInvalidExpireAtTime = errors.New("invalid expire at time")
	ErrInvalidUserPercent  = errors.New("invalid user percent")
	ErrNotEnoughUsers      = errors.New("not enough users in db")
)
