package usecase

import "github.com/Tomoki108/burny/domain"

type APIKeyUseCase struct {
	repo          domain.APIKeyRepository
	transactioner domain.Transactioner
}

func (u *APIKeyUseCase) Get(userID uint) (*domain.APIKey, error) {
	return u.repo.GetByUserID(u.transactioner.Default(), userID)
}

func (u *APIKeyUseCase) Create(userID uint) (*domain.APIKey, error) {
	key := &domain.APIKey{
		UserID: userID,
	}
	return u.repo.Create(u.transactioner.Default(), key)
}

func (u *APIKeyUseCase) Delete(userID uint) error {
	return u.repo.DeleteByUserID(u.transactioner.Default(), userID)
}
