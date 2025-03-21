package usecase

import "github.com/Tomoki108/burny/domain"

type APIKeyUseCase struct {
}

func (u *APIKeyUseCase) Get(userID uint) (*domain.APIKey, error) {
	return nil, nil
}

func (u *APIKeyUseCase) Create(userID uint) (*domain.APIKey, error) {
	return nil, nil
}

func (u *APIKeyUseCase) Delete(userID uint) error {
	return nil
}
