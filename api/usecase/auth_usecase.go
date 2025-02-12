package usecase

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/Tomoki108/burny/config"
	"github.com/Tomoki108/burny/domain"
	"golang.org/x/crypto/bcrypt"
)

var ErrEmailAlreadyExists = errors.New("メールアドレスが既に使用されています")
var ErrUserNotExists = errors.New("ユーザーが存在しません")
var ErrInvalidPassword = errors.New("パスワードが間違っています")

type AuthUseCase struct {
	Repo          domain.UserRepository
	Transactioner domain.Transactioner
}
type JwtCustomClaims struct {
	Email string
	jwt.RegisteredClaims
}

func (u AuthUseCase) SignUp(user *domain.User) error {
	exisitingUser, err := u.Repo.GetByEmail(u.Transactioner.Default(), user.Email)
	if err != nil {
		return err
	}
	if exisitingUser != nil {
		return ErrEmailAlreadyExists
	}

	_, err = u.Repo.Create(u.Transactioner.Default(), user)
	return err
}

func (u AuthUseCase) SignIn(user *domain.User) (tokenStr string, err error) {
	exisitingUser, err := u.Repo.GetByEmail(u.Transactioner.Default(), user.Email)
	if err != nil {
		return "", err
	}
	if exisitingUser == nil {
		return "", ErrUserNotExists
	}

	if err := bcrypt.CompareHashAndPassword([]byte(exisitingUser.Password), []byte(user.Password)); err != nil {
		return "", ErrInvalidPassword
	}

	claims := &JwtCustomClaims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Conf.JwtSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}
