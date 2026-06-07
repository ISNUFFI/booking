package errs

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrDuplicateKey       = errors.New("duplicate key")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidPassword    = errors.New("invalid password")
)
