package domain

import (
	"time"
)

type User struct {
	ID            uint      `json:"id"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"email_verified"`
	Password      string    `json:"-"` // always must be hashed
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(tx Transaction, user *User) (*User, error)
	Get(tx Transaction, id uint) (*User, error)
	GetByEmail(tx Transaction, email string) (*User, error)
	Update(tx Transaction, user *User) (*User, error)
}

const UserEmailVerifiedTopic = "user:email:verified"

type UserEmailVerifiedEvent struct {
	UserID uint
}
