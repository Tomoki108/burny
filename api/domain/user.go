package domain

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NOTE: 取得したUserのパスワードはハッシュ化されている
type UserRepository interface {
	// passwordのハッシュ化はinfrastructure層で行う
	Create(user *User) (*User, error)
	Get(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
}
