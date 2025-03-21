package model

import (
	"github.com/Tomoki108/burny/domain"
	"gorm.io/gorm"
)

type APIKey struct {
	gorm.Model
	UserID uint   `json:"user_id" gorm:"index"`
	Key    string `json:"key" gorm:"uniqueIndex"`
}

func (a *APIKey) ToDomain() *domain.APIKey {
	return &domain.APIKey{
		ID:        a.ID,
		UserID:    a.UserID,
		Key:       a.Key,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func FromDomainAPIKey(apiKey *domain.APIKey) *APIKey {
	return &APIKey{
		Model: gorm.Model{
			ID:        apiKey.ID,
			CreatedAt: apiKey.CreatedAt,
			UpdatedAt: apiKey.UpdatedAt,
		},
		UserID: apiKey.UserID,
		Key:    apiKey.Key,
	}
}
