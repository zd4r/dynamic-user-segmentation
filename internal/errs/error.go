package errs

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrInvalidUserId = errors.New("invalid user id")
)
