package infrastructure

import (
	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/model"
)

func NewUserRepository() domain.UserRepository {
	return &UserRepository{}
}

type UserRepository struct {
}

func (r *UserRepository) Create(user *domain.User) (*domain.User, error) {
	model, err := model.FromDomainUser(user)
	if err != nil {
		return nil, err
	}
	if err := DB.Create(model).Error; err != nil {
		return nil, err
	}

	return r.Get(model.ID)
}

func (r *UserRepository) Get(id uint) (*domain.User, error) {
	var user model.User
	if err := DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user.ToDomain(), nil
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var user model.User
	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user.ToDomain(), nil
}
