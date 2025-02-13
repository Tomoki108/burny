package infrastructure

import (
	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/model"
	"gorm.io/gorm"
)

func NewSprintRepository() domain.SprintRepository {
	return &SprintRepository{}
}

type SprintRepository struct {
}

func (r *SprintRepository) Create(tx domain.Transaction, sprint *domain.Sprint) (*domain.Sprint, error) {
	db := tx.(*gorm.DB)
	model := model.FromDomainSprint(sprint)
	if err := db.Create(model).Error; err != nil {
		return nil, err
	}

	return r.Get(tx, model.ID)
}

func (r *SprintRepository) List(tx domain.Transaction) ([]*domain.Sprint, error) {
	db := tx.(*gorm.DB)
	var sprints []model.Sprint
	if err := db.Find(&sprints).Error; err != nil {
		return nil, err
	}

	domains := make([]*domain.Sprint, len(sprints))
	for i, sprint := range sprints {
		domains[i] = sprint.ToDomain()
	}

	return domains, nil
}

func (r *SprintRepository) Get(tx domain.Transaction, id uint) (*domain.Sprint, error) {
	db := tx.(*gorm.DB)
	var sprint model.Sprint
	if err := db.First(&sprint, id).Error; err != nil {
		return nil, err
	}

	return sprint.ToDomain(), nil
}

func (r *SprintRepository) Update(tx domain.Transaction, sprint *domain.Sprint) (*domain.Sprint, error) {
	db := tx.(*gorm.DB)
	model := model.FromDomainSprint(sprint)
	if err := db.Save(model).Error; err != nil {
		return nil, err
	}

	return r.Get(tx, model.ID)
}
