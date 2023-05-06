package models

import "time"

type Bill struct {
	ID            string    `json:"id"`
	From          time.Time `json:"from"`
	Until         time.Time `json:"until"`
	SpendAmount   float64   `json:"spend_amount"`
	PaymentAmount float64   `json:"payment_amount"`
	ApartmentId   string    `json:"apartment_id"`
	ServiceId     string    `json:"service_id"`
	Paid          bool      `json:"paid"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
