package handler

import (
	"errors"
	"net/http"

	"github.com/Tomoki108/burny/domain"
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
// @Success      201
// @Failure      400
// @Failure      500
// @Router       /sign_up [post]
func (h AuthHandler) SignUp(c echo.Context) error {
	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
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
// @Success      200
// @Failure      400
// @Failure      500
// @Router        /sign_in [post]
func (h AuthHandler) SignIn(c echo.Context) error {
	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return err
}
