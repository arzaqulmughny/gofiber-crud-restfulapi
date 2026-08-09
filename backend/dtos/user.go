package dtos

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
