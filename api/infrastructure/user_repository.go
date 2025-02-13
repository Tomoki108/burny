package infrastructure

import (
	"errors"

	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/model"
	"gorm.io/gorm"
)

func NewUserRepository() domain.UserRepository {
	return &UserRepository{}
}

type UserRepository struct {
}

func (r *UserRepository) Create(tx domain.Transaction, user *domain.User) (*domain.User, error) {
	model, err := model.FromDomainUser(user)
	if err != nil {
		return nil, err
	}
	if err := DB.Create(model).Error; err != nil {
		return nil, err
	}

	return r.Get(tx, model.ID)
}

func (r *UserRepository) Get(tx domain.Transaction, id uint) (*domain.User, error) {
	var user model.User
	if err := DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user.ToDomain(), nil
}

func (r *UserRepository) GetByEmail(tx domain.Transaction, email string) (*domain.User, error) {
	var user model.User

	err := DB.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return user.ToDomain(), nil
}
