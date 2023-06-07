package models

type ResponseError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type JsonResponse struct {
	Error *ResponseError `json:"error"`
	Data  interface{}    `json:"data"`
}
