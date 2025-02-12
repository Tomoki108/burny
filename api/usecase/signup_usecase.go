package usecase

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/Tomoki108/burny/domain"
	"golang.org/x/crypto/bcrypt"
)

var ErrEmailAlreadyExists = errors.New("メールアドレスが既に使用されています")
var ErrUserNotExists = errors.New("ユーザーが存在しません")
var ErrInvalidPassword = errors.New("パスワードが間違っています")

type SignUpUseCase struct {
	Repo domain.UserRepository
}
type JwtCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
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

func (u SignUpUseCase) SignIn(user *domain.User) (tokenStr string, err error) {
	exisitingUser, err := u.Repo.GetByEmail(user.Email)
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
	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return t, nil
}
