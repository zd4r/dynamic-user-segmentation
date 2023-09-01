package errs

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrInvalidUserId       = errors.New("invalid user id")
	ErrInvalidExpireAtTime = errors.New("invalid expire at time")
	ErrInvalidUserPercent  = errors.New("invalid user percent")
	ErrNotEnoughUsers      = errors.New("not enough users in db")
)
