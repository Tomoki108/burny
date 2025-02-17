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

func NewAuthHandler(usecase usecase.AuthUseCase) AuthHandler {
	return AuthHandler{
		Usecase: usecase,
	}
}

// @Summary      Sign up
// @Description  Sign up
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body io.SignUpRequest true "sign up request"
// @Success      201 {object} domain.User
// @Failure      400 {object} io.ValidationErrorResponse
// @Router       /sign_up [post]
func (h AuthHandler) SignUp(c echo.Context) error {
	req := new(io.SignUpRequest)
	if err := handleReq(c, req); err != nil {
		return err
	}

	user, err := h.Usecase.SignUp(*req)
	if errors.Is(err, usecase.ErrEmailAlreadyExists) {
		return c.JSON(http.StatusConflict, io.NewErrResp(err.Error()))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, io.NewErrResp(err.Error()))
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
// @Failure      400 {object} io.ValidationErrorResponse
// @Router       /sign_in [post]
func (h AuthHandler) SignIn(c echo.Context) error {
	req := new(io.SignInRequest)
	if err := handleReq(c, req); err != nil {
		return err
	}

	jwtToken, err := h.Usecase.SignIn(*req)
	if errors.Is(err, usecase.ErrUserNotExists) || errors.Is(err, usecase.ErrInvalidPassword) {
		return c.JSON(http.StatusUnauthorized, io.NewErrResp(err.Error()))
	}

	res := io.SignInResponse{
		JwtToken: jwtToken,
	}
	return c.JSON(http.StatusOK, res)
}
