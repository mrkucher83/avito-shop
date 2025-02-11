package models

type Employee struct {
	Id       int    `json:"-"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Coins    int    `json:"coins"`
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string
}
