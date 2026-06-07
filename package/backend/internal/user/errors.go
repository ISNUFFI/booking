package user

import "errors"

var ErrDuplicateKey = errors.New("duplicate key")
var ErrEmailAlreadyExists = errors.New("email already exists")
var ErrInvalidPassword = errors.New("invalid password")
