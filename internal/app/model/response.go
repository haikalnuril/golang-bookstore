package model

type Response struct {
	Code int `json:"code"` // HTTP status code
	Message string `json:"message"` // Response message
	Data any `json:"data"`
}
