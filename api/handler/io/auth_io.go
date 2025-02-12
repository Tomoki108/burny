package io

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInRequest SignUpRequest

type SignInResponse struct {
	JwtToken string `json:"token"`
}
