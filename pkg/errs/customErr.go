package errs

import "errors"

var (
	ErrUserEmailExists    = errors.New("email already exists")
	ErrUserUsernameExists = errors.New("username already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrGameNotFound       = errors.New("game not found")
)
