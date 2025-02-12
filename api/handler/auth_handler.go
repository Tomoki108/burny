package handler

import "github.com/Tomoki108/burny/domain"

type AuthHandler struct {
	Repo domain.UserRepository
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
	_, err := h.Repo.Create(user)

	return err
}
