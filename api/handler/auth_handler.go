package handler

import (
	"errors"
	"net/http"

	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/handler/io"
	"github.com/Tomoki108/burny/usecase"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Usecase usecase.SignUpUseCase
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
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user := &domain.User{
		Email:    req.Email,
		Password: req.Password,
	}
	err := h.Usecase.SignUp(user)
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
	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	jwtToken, err := h.Usecase.SignIn(user)
	if errors.Is(err, usecase.ErrUserNotExists) || errors.Is(err, usecase.ErrInvalidPassword) {
		return c.JSON(http.StatusUnauthorized, err)
	}

	res := io.SignInResponse{
		JwtToken: jwtToken,
	}
	return c.JSON(http.StatusOK, res)
}
