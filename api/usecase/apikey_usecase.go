package usecase

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/handler/io"
	"golang.org/x/crypto/bcrypt"
)

var ErrAPIKeyAlreadyExists = errors.New("API Key already exists, please delete it first")
var ErrAPIKeyNotFound = errors.New("API Key not exists")

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

func (u *APIKeyUseCase) CheckStatus(userID uint) (bool, error) {
	apiKey, err := u.repo.GetByUserID(u.transactioner.Default(), userID)
	if err != nil {
		return false, err
	}
	return apiKey != nil, nil
}

func (u *APIKeyUseCase) Create(userID uint) (*io.CreateAPIKeyResponse, error) {
	exists, err := u.CheckStatus(userID)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, ErrAPIKeyAlreadyExists
	}

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
	_, err = u.repo.Create(u.transactioner.Default(), key)
	if err != nil {
		return nil, err
	}

	return &io.CreateAPIKeyResponse{
		RawKey: rawKey,
	}, nil
}

func (u *APIKeyUseCase) Delete(userID uint) error {
	exists, err := u.CheckStatus(userID)
	if err != nil {
		return err
	}

	if !exists {
		return ErrAPIKeyNotFound
	}

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
