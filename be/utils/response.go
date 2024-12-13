package utils

type MetaData struct {
	TotalPage int   `json:"totalPage"`
	TotalData int64 `json:"totalData"`
	Page      int   `json:"page"`
}
type GlobalResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    MetaData    `json:"meta"`
}
