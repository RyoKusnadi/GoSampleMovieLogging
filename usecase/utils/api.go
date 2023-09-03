package utils

import "scalable-go-movie/domain"

func CreateSuccessApiResponse(result interface{}) domain.ApiResponse {
	return domain.ApiResponse{
		Result: result,
	}
}

func CreateDefaultErrorApiResponse(code int, message string) domain.ApiResponse {
	return domain.ApiResponse{
		Error: &domain.ErrorResponse{
			Code:    code,
			Message: message,
		},
	}
}
