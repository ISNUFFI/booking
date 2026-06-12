package bookings

import (
	"time"
)

type Booking struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id,omitempty"`
	SlotID    int       `json:"slot_id"`
	CreatedAt time.Time `json:"created_at"`
}
