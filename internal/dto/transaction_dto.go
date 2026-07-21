package dto

import "time"

type BuyRequest struct {
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

type SellRequest struct {
	Gram float64 `json:"gram" binding:"required,gt=0"`
}

type PriceResponse struct {
	Price float64 `json:"price"`
}

type BuyResponse struct {
	Gram  float64 `json:"gram"`
	Price float64 `json:"price"`
}

type SellResponse struct {
	Amount float64 `json:"amount"`
}

type TransactionResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Type      string    `json:"type"`
	Amount    float64   `json:"amount"`
	Gram      float64   `json:"gram"`
	CreatedAt time.Time `json:"created_at"`
}
