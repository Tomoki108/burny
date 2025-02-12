package usecase

import (
	"errors"

	"github.com/Tomoki108/burny/domain"
)

var ErrEmailAlreadyExists = errors.New("email already exists")

type SignUpUseCase struct {
	Repo domain.UserRepository
}

func (u SignUpUseCase) SignUp(user *domain.User) error {
	exisitingUser, err := u.Repo.GetByEmail(user.Email)
	if err != nil {
		return err
	}
	if exisitingUser != nil {
		return ErrEmailAlreadyExists
	}

	_, err = u.Repo.Create(user)
	return err
}
