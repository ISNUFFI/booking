package slots

import (
	"time"
)

type Slot struct {
	ID         int       `json:"id"`
	ProviderID int       `json:"provider_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	IsActive   bool      `json:"is_active"`
}
