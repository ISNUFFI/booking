package users

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	passwordHash string
}
