package response

import "person-service/internal/core/entity/error_code"

type Response struct {
	Data         interface{}          `json:"data"`
	Status       bool                 `json:"status"`
	ErrorCode    error_code.ErrorCode `json:"errorCode"`
	ErrorMessage string               `json:"errorMessage"`
}

type SignUpDataResponse struct {
	DisplayName string `json:"displayName"`
}
