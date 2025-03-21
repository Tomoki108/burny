package infrastructure

import (
	"errors"

	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/model"
	"gorm.io/gorm"
)

func NewAPIKeyRepository() domain.APIKeyRepository {
	return &APIKeyRepository{}
}

type APIKeyRepository struct {
}

func (r *APIKeyRepository) Create(tx domain.Transaction, apiKey *domain.APIKey) (*domain.APIKey, error) {
	db := tx.(*gorm.DB)
	model := model.FromDomainAPIKey(apiKey)
	if err := db.Create(model).Error; err != nil {
		return nil, err
	}
	return model.ToDomain(), nil
}

func (r *APIKeyRepository) GetByUserID(tx domain.Transaction, userID uint) (*domain.APIKey, error) {
	db := tx.(*gorm.DB)
	var apiKey model.APIKey
	err := db.Where("user_id = ?", userID).First(&apiKey).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return apiKey.ToDomain(), nil
}

func (r *APIKeyRepository) DeleteByUserID(tx domain.Transaction, userID uint) error {
	db := tx.(*gorm.DB)
	return db.Where("user_id = ?", userID).Delete(&model.APIKey{}).Error
}

// GetAll すべてのAPIキーを取得する
func (r *APIKeyRepository) GetAll(tx domain.Transaction) ([]*domain.APIKey, error) {
	var db *gorm.DB
	if tx == nil {
		db = DB
	} else {
		db = tx.(*gorm.DB)
	}

	var apiKeys []model.APIKey
	if err := db.Find(&apiKeys).Error; err != nil {
		return nil, err
	}

	// モデルをドメインオブジェクトに変換
	result := make([]*domain.APIKey, len(apiKeys))
	for i, key := range apiKeys {
		result[i] = key.ToDomain()
	}

	return result, nil
}
