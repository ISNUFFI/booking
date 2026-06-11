package slots

import "errors"

var (
	ErrSlotNotFound = errors.New("slot not found")
	ErrStartAfterEnd = errors.New("start time should be before end time")
	ErrDurationNotPositive = errors.New("duration should be greater than 0")
)
