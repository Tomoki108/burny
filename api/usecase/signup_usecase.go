package usecase

import (
	"errors"

	"github.com/Tomoki108/burny/domain"
	"golang.org/x/crypto/bcrypt"
)

var ErrEmailAlreadyExists = errors.New("メールアドレスが既に使用されています")
var ErrUserNotExists = errors.New("ユーザーが存在しません")
var ErrInvalidPassword = errors.New("パスワードが間違っています")

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

func (u SignUpUseCase) SignIn(user *domain.User) error {
	exisitingUser, err := u.Repo.GetByEmail(user.Email)
	if err != nil {
		return err
	}
	if exisitingUser == nil {
		return ErrUserNotExists
	}

	if err := bcrypt.CompareHashAndPassword([]byte(exisitingUser.Password), []byte(user.Password)); err != nil {
		return ErrInvalidPassword
	}

	_, err = u.Repo.Create(user)
	return err
}
