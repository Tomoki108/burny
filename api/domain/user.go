package domain

import (
	"encoding/json"
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"` // always must be hashed
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MarshalJSONをカスタマイズし、パスワードをレスポンスから取り除く
func (u User) MarshalJSON() ([]byte, error) {
	type Alias User // エイリアスを作ることで、無限再帰を回避
	return json.Marshal(&struct {
		Alias
		Password string `json:"-"`
	}{
		Alias: Alias(u),
	})
}

type UserRepository interface {
	Create(tx Transaction, user *User) (*User, error)
	Get(tx Transaction, id uint) (*User, error)
	GetByEmail(tx Transaction, email string) (*User, error)
}
