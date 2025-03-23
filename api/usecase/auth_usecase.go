package usecase

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/Tomoki108/burny/config"
	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/handler/io"
	evbus "github.com/asaskevich/EventBus"
	"golang.org/x/crypto/bcrypt"
)

var ErrEmailAlreadyExists = errors.New("email has already been registered")
var ErrUserNotExists = errors.New("user not exists")
var ErrInvalidPassword = errors.New("password is invalid")

type AuthUseCase struct {
	Repo          domain.UserRepository
	Transactioner domain.Transactioner
	EventBus      evbus.Bus
	Mailer        domain.Mailer // TODO: mod constructor and DI setting
}

func NewAuthUseCase(repo domain.UserRepository, transactioner domain.Transactioner, eventBus evbus.Bus, mailer domain.Mailer) AuthUseCase {
	return AuthUseCase{
		Repo:          repo,
		Transactioner: transactioner,
		EventBus:      eventBus,
		Mailer:        mailer,
	}
}

func (u AuthUseCase) SignUp(req io.SignUpRequest) (*domain.User, error) {
	exisitingUser, err := u.Repo.GetByEmail(u.Transactioner.Default(), req.Email)
	if err != nil {
		return nil, err
	}
	if exisitingUser != nil {
		return nil, ErrEmailAlreadyExists
	}

	// NOTE: ハッシュ生成時には内部でランダムなソルトが利用される。ハッシュにはソルトの情報も含まれるため、後で平文と比較検証ができる。
	hassedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		Email:    req.Email,
		Password: string(hassedPassword),
	}

	err = u.Transactioner.Transaction(func(tx domain.Transaction) error {
		user, err = u.Repo.Create(tx, user)
		if err != nil {
			return err
		}

		claims := &jwt.MapClaims{
			"user_id": user.ID,
			"email":   user.Email,
			"exp":     time.Now().Add(time.Minute * 20).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		verificationToken, err := token.SignedString([]byte(config.Conf.JwtSecret))
		if err != nil {
			return err
		}

		mail := domain.NewEmailVerificationMail(user.Email, verificationToken)
		if err := u.Mailer.Send(mail); err != nil {
			return err
		}

		u.EventBus.Publish(domain.UserCreatedTopic, domain.UserCreatedEvent{UserID: user.ID})
		return nil
	})

	if err != nil {
		return nil, err
	}
	return user, nil
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
