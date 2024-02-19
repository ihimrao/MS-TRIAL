package utils

type ApiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(statusCode int, message string, data interface{}) ApiResponse {
	return ApiResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(statusCode int, message string) ApiResponse {
	return ApiResponse{
		Status:  statusCode,
		Message: message,
	}
}
