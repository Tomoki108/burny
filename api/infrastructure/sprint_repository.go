package infrastructure

import (
	"errors"

	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/model"
	"gorm.io/gorm"
)

var ErrrSprintNotFound = errors.New("スプリントが存在しません")

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

	return r.Get(tx, model.ID, model.ProjectID)
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
	db := tx.(*gorm.DB)
	var sprint model.Sprint
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
		return nil, ErrrSprintNotFound
	}
	if err != nil {
		return nil, err
	}

	sprint.ActualSP = actualSP
	if err := db.Save(sprint).Error; err != nil {
		return nil, err
	}

	return r.Get(tx, sprint.ID, sprint.ProjectID)
}
