package usecase

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/Tomoki108/burny/domain"
	"golang.org/x/crypto/bcrypt"
)

type APIKeyUseCase struct {
	repo          domain.APIKeyRepository
	transactioner domain.Transactioner
}

func NewAPIKeyUseCase(repo domain.APIKeyRepository, transactioner domain.Transactioner) APIKeyUseCase {
	return APIKeyUseCase{
		repo:          repo,
		transactioner: transactioner,
	}
}

func (u *APIKeyUseCase) Get(userID uint) (*domain.APIKey, error) {
	return u.repo.GetByUserID(u.transactioner.Default(), userID)
}

func (u *APIKeyUseCase) Create(userID uint) (*domain.APIKey, error) {
	// Generate a random API key
	rawKey, err := generateRandomKey(32) // 32 bytes = 256 bits
	if err != nil {
		return nil, fmt.Errorf("failed to generate API key: %w", err)
	}

	// Hash the API key for storage
	hashedKey, err := bcrypt.GenerateFromPassword([]byte(rawKey), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash API key: %w", err)
	}

	key := &domain.APIKey{
		UserID: userID,
		Key:    string(hashedKey),
	}

	// Store the hashed key in the database
	createdKey, err := u.repo.Create(u.transactioner.Default(), key)
	if err != nil {
		return nil, err
	}

	// Return the created key with the raw (unhashed) key for the client
	// This is the only time the raw key will be exposed
	createdKey.Key = rawKey
	return createdKey, nil
}

func (u *APIKeyUseCase) Delete(userID uint) error {
	return u.repo.DeleteByUserID(u.transactioner.Default(), userID)
}

// generateRandomKey creates a cryptographically secure random key with the specified byte length
// and returns it as a base64-encoded string
func generateRandomKey(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
