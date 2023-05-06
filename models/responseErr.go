package models

type ResponseError struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
}
