package model

import (
	"github.com/Tomoki108/burny/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email         string `json:"email" gorm:"uniqueIndex"`
	EmailVerified bool   `json:"email_verified"`
	Password      string `json:"password"`
}

func (u *User) ToDomain() *domain.User {
	return &domain.User{
		ID:            u.ID,
		Email:         u.Email,
		EmailVerified: u.EmailVerified,
		Password:      u.Password,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}
}

func FromDomainUser(user *domain.User) *User {
	return &User{
		Model: gorm.Model{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		Password:      user.Password,
	}
}
