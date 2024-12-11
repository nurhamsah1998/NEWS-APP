package utils

type GlobalResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
