package dtos

import "gofiber-restful-api/domain"

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Message string `json="message"`
	Data    string `json="data"`
}

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRegisterResponse struct {
	Message string       `json:"message"`
	Data    domain.Users `json:"data"`
}
