package utils

type Response struct {
	message string
	data    interface{}
}

func Res(message string, data interface{}) Response {
	return Response{message: message, data: data}
}
