package models

// Employee models -------------------------

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Employee struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Coins    int    `json:"coins"`
}

type AuthResponse struct {
	Token string
}

// Purchase models -------------------------

type Merch struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// SendCoin models -------------------------

type SendCoinRequest struct {
	ToUser string `json:"toUser"`
	Amount int    `json:"amount"`
}

type Response struct {
	Message string `json:"message"`
}
