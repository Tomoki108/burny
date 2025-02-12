package model

import (
	"github.com/Tomoki108/burny/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"`
}

func (u *User) ToDomain() *domain.User {
	return &domain.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}
}

func FromDomainUser(user *domain.User) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:    user.Email,
		Password: string(hashedPassword),
	}, nil
}
