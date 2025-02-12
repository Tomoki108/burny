package handler

import (
	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/usecase"
)

type AuthHandler struct {
	Usecase usecase.SignUpUseCase
}

// @Summary      Sign up
// @Description  Sign up
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /projects [get]
func (h AuthHandler) SignUp(user *domain.User) error {
	err := h.Usecase.SignUp(user)

	return err
}
