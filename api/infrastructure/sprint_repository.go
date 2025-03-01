package infrastructure

import (
	"errors"

	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/model"
	"github.com/Tomoki108/burny/usecase"
	"gorm.io/gorm"
)

func NewSprintRepository() domain.SprintRepository {
	return &SprintRepository{}
}

type SprintRepository struct {
}

func (r *SprintRepository) Create(tx domain.Transaction, sprint *domain.Sprint) error {
	db := tx.(*gorm.DB)
	model := model.FromDomainSprint(sprint)

	return db.Create(model).Error
}

func (r *SprintRepository) List(tx domain.Transaction, pojectID uint) ([]*domain.Sprint, error) {
	db := tx.(*gorm.DB)
	var sprints []model.Sprint
	if err := db.Where("project_id = ?", pojectID).Order("start_date ASC").Find(&sprints).Error; err != nil {
		return nil, err
	}

	domains := make([]*domain.Sprint, len(sprints))
	for i, sprint := range sprints {
		domains[i] = sprint.ToDomain()
	}

	return domains, nil
}

func (r *SprintRepository) Get(tx domain.Transaction, pojectID, sprintID uint) (*domain.Sprint, error) {
	var sprint model.Sprint
	db := tx.(*gorm.DB)
	if err := db.Where("id = ? AND project_id = ?", sprintID, pojectID).First(&sprint).Error; err != nil {
		return nil, err
	}

	return sprint.ToDomain(), nil
}

func (r *SprintRepository) Update(tx domain.Transaction, projectID, sprintID uint, actualSP int) (*domain.Sprint, error) {
	var sprint model.Sprint

	db := tx.(*gorm.DB)
	err := db.Where("id = ? AND project_id = ?", sprintID, projectID).First(&sprint).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, usecase.ErrSprintNotFound
	}
	if err != nil {
		return nil, err
	}

	sprint.ActualSP = actualSP
	if err := db.Save(sprint).Error; err != nil {
		return nil, err
	}

	return r.Get(tx, sprint.ProjectID, sprint.ID)
}

func (r *SprintRepository) Delete(tx domain.Transaction, projectID, sprintID uint) error {
	db := tx.(*gorm.DB)
	if err := db.Where("id = ? AND project_id = ?", sprintID, projectID).Delete(&model.Sprint{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *SprintRepository) UpsertList(tx domain.Transaction, sprints []*domain.Sprint) ([]*domain.Sprint, error) {
	db := tx.(*gorm.DB)
	models := make([]model.Sprint, 0, len(sprints))
	for _, sprint := range sprints {
		model := model.FromDomainSprint(sprint)
		models = append(models, *model)
		if err := db.Save(model).Error; err != nil {
			return nil, err
		}
	}

	domains := make([]*domain.Sprint, 0, len(models))
	for _, model := range models {
		domains = append(domains, model.ToDomain())
	}
	return domains, nil
}
