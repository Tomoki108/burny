package io

type SignUpRequest struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

type SignInRequest SignUpRequest

type SignInResponse struct {
	JwtToken string `json:"token"`
}

type VerifyEmailRequest struct {
	Token string `json:"-" param:"token"`
}
