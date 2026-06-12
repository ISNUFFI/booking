package bookings

import (
	"time"
)

type Booking struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id,omitempty"`
	SlotID    int64     `json:"slot_id"`
	CreatedAt time.Time `json:"created_at"`
}
