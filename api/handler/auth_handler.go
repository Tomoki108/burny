package handler

import (
	"errors"
	"net/http"

	"github.com/Tomoki108/burny/handler/io"
	"github.com/Tomoki108/burny/usecase"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Usecase usecase.AuthUseCase
}

// @Summary      Sign up
// @Description  Sign up
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body io.SignUpRequest true "sign up request"
// @Success      201 {object} domain.User
// @Failure      400
// @Failure      500
// @Router       /sign_up [post]
func (h AuthHandler) SignUp(c echo.Context) error {
	req := new(io.SignUpRequest)
	if err := handleReq(c, req); err != nil {
		return err
	}

	user, err := h.Usecase.SignUp(*req)
	if errors.Is(err, usecase.ErrEmailAlreadyExists) {
		return c.JSON(http.StatusConflict, err)
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, user)
}

// @Summary      Sign in
// @Description  Sign in
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body io.SignInRequest true "sign in request"
// @Success      200 {object} io.SignInResponse
// @Failure      400
// @Failure      500
// @Router       /sign_in [post]
func (h AuthHandler) SignIn(c echo.Context) error {
	req := new(io.SignInRequest)
	if err := handleReq(c, req); err != nil {
		return err
	}

	jwtToken, err := h.Usecase.SignIn(*req)
	if errors.Is(err, usecase.ErrUserNotExists) || errors.Is(err, usecase.ErrInvalidPassword) {
		return c.JSON(http.StatusUnauthorized, err)
	}

	res := io.SignInResponse{
		JwtToken: jwtToken,
	}
	return c.JSON(http.StatusOK, res)
}
