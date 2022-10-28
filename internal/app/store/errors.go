package store

import "errors"

var (
	ErrUserNotFound error = errors.New("user not found")
)
