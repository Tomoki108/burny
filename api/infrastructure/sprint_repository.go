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

	return r.Get(model.ID)
}

func (r *SprintRepository) List() ([]*domain.Sprint, error) {
	var sprints []model.Sprint
	if err := DB.Find(&sprints).Error; err != nil {
		return nil, err
	}

	domains := make([]*domain.Sprint, len(sprints))
	for i, sprint := range sprints {
		domains[i] = sprint.ToDomain()
	}

	return domains, nil
}

func (r *SprintRepository) Get(id uint) (*domain.Sprint, error) {
	var sprint model.Sprint
	if err := DB.First(&sprint, id).Error; err != nil {
		return nil, err
	}

	return sprint.ToDomain(), nil
}

func (r *SprintRepository) Update(sprint *domain.Sprint) (*domain.Sprint, error) {
	model := model.FromDomainSprint(sprint)
	if err := DB.Save(model).Error; err != nil {
		return nil, err
	}

	return r.Get(model.ID)
}
