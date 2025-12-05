package errs

import "errors"

var (
	ErrUserEmailExists    = errors.New("email already exists")
	ErrUserUsernameExists = errors.New("username already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrInternal           = errors.New("internal server error")
)
