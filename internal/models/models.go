package models

// Employee models ------------------------------------

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

// Purchase models ------------------------------------

type Merch struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Coins models -----------------------------------

type SendCoin struct {
	ToUser string `json:"toUser"`
	Amount int    `json:"amount"`
}

type ReceiveCoin struct {
	FromUser string `json:"fromUser"`
	Amount   int    `json:"amount"`
}

// Get employee info models --------------------------

type EmployeeInfoResponse struct {
	Coins        int `json:"coins"`
	Inventory    `json:"inventory"`
	CoinsHistory `json:"coinsHistory"`
}

type Inventory struct {
	Items []Purchase `json:"items"`
}

type Purchase struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type Received struct {
	Items []ReceiveCoin `json:"items"`
}

type Sent struct {
	Items []SendCoin `json:"items"`
}

type CoinsHistory struct {
	Received `json:"received"`
	Sent     `json:"sent"`
}

// Standard response message ------------------------

type Response struct {
	Message string `json:"message"`
}
