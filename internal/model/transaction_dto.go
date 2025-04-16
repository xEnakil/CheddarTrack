package model

import "time"

type CreateTransactionInput struct {
	UserId      uint      `json:"user_id" binding:"required"`
	CategoryID  uint      `json:"category_id" binding:"required"`
	Amount      float64   `json:"amount" binding:"required"`
	Currency    string    `json:"currency" binding:"required"`
	Description string    `json:"description"`
	Timestamp   time.Time `json:"timestamp" binding:"required"`
}

type TransactionResponse struct {
	ID          uint      `json:"id"`
	Amount      float64   `json:"amount"`
	Currency    string    `json:"currency"`
	Description string    `json:"description"`
	Timestamp   time.Time `json:"timestamp"`
	Category    string    `json:"category"`
	Type        string    `json:"type"`
}
