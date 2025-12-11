package errs

import "errors"

var (
	ErrUserEmailExists    = errors.New("email already exists")
	ErrUserUsernameExists = errors.New("username already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrGameNotFound       = errors.New("game not found")
	ErrUserPassword       = errors.New("invalid password and/or login")
	ErrReviewExists       = errors.New("already reviewed")
	ErrReviesNotFound     = errors.New("reviews not found")
	ErrInvalidMetadata    = errors.New("invalid metadata")
	ErrMetadataNotFound   = errors.New("metadata not found")
)
