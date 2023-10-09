package service

import (
	"person-service/internal/core/model/request"
	"person-service/internal/core/model/response"
)

type UserService interface {
	SignUp(request *request.SignUpRequest) *response.Response
}
