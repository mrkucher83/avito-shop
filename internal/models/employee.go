package models

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Employee struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string
}
