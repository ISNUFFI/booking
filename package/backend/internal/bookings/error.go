package bookings

import (
	"errors"
)

var (
	ErrSlotAlreadyTaken = errors.New("slot already taken")
	ErrSlotDoesNotExist = errors.New("slot already taken")
	ErrForbidden        = errors.New("forbidden")
	ErrBookingNotFound  = errors.New("booking not found")
)
