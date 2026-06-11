package providers

import "errors"

var (
	ErrProviderNotFound      = errors.New("provider not found")
	ErrProviderOwnerMismatch = errors.New("user does not own the provider")
)
