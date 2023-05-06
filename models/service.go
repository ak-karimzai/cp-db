package models

type Service struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	MAmount     string  `json:"m_amount"`
	Cost        float64 `json:"cost"`
}
