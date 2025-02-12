package domain

import "encoding/json"

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
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

// NOTE: 取得したUserのパスワードはハッシュ化されている
type UserRepository interface {
	// passwordのハッシュ化はinfrastructure層で行う
	Create(user *User) (*User, error)
	Get(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
}
