package models

import "time"

type Payment struct {
	ID        string    `json:"id"`
	Amount    float64   `json:"amount"`
	BillId    string    `json:"bill_id"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
