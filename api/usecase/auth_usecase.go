package usecase

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/Tomoki108/burny/config"
	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/handler/io"
	"golang.org/x/crypto/bcrypt"
)

var ErrEmailAlreadyExists = errors.New("email has already been registered")
var ErrUserNotExists = errors.New("user not exists")
var ErrInvalidPassword = errors.New("password is invalid")

type AuthUseCase struct {
	Repo          domain.UserRepository
	Transactioner domain.Transactioner
}

func (u AuthUseCase) SignUp(req io.SignUpRequest) (*domain.User, error) {
	exisitingUser, err := u.Repo.GetByEmail(u.Transactioner.Default(), req.Email)
	if err != nil {
		return nil, err
	}
	if exisitingUser != nil {
		return nil, ErrEmailAlreadyExists
	}

	hassedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		Email:    req.Email,
		Password: string(hassedPassword),
	}

	return u.Repo.Create(u.Transactioner.Default(), user)
}

func (u AuthUseCase) SignIn(req io.SignInRequest) (tokenStr string, err error) {
	user, err := u.Repo.GetByEmail(u.Transactioner.Default(), req.Email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", ErrUserNotExists
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", ErrInvalidPassword
	}

	claims := &jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Conf.JwtSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}
