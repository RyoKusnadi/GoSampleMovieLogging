package domain

type Response struct {
	Result interface{}    `json:"result,omitempty"`
	Error  *ErrorResponse `json:"error,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func CreateSuccessResponse(result interface{}) Response {
	return Response{
		Result: result,
	}
}

func CreateErrorResponse(code int, message string) Response {
	return Response{
		Error: &ErrorResponse{
			Code:    code,
			Message: message,
		},
	}
}
