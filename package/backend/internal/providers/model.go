package providers

type Provider struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       int    `json:"owner_id"`
}
