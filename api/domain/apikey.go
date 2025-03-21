package domain

import (
	"time"
)

// APIKey represents an API key for authentication in the domain layer
type APIKey struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// APIKeyRepository defines the interface for API key storage
type APIKeyRepository interface {
	Create(tx Transaction, apiKey *APIKey) (*APIKey, error)
	GetByUserID(tx Transaction, userID uint) (*APIKey, error)
	DeleteByUserID(tx Transaction, userID uint) error
	GetAll(tx Transaction) ([]*APIKey, error) // 追加：すべてのAPIキーを取得するメソッド
}
