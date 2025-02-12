package model

import (
	"github.com/Tomoki108/burny/domain"
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

func FromDomainUser(user *domain.User) *User {
	return &User{
		Email:    user.Email,
		Password: user.Password,
	}
}
