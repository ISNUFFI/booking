package slots

type Slot struct {
	ID			int		`json:"id"`
	ProviderID  string  `json:"provider_id"`
	StartTime  	string  `json:"start_time"`
	EndTime		int	    `json:"end_time"`
	IsActive	bool	`json:"is_active"`
}
