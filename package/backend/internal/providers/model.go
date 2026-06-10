package providers

type Provider struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner		int	   `json:"owner_id"`
}
