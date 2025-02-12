package handler

import "github.com/Tomoki108/burny/domain"

type UserHandler struct {
	Repo domain.UserRepository
}

func (h UserHandler) SignUp(user *domain.User) (*domain.User, error) {
	return h.Repo.Create(user)
}
