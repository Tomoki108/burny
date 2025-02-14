package infrastructure

import (
	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/model"
	"gorm.io/gorm"
)

func NewProjectRepository() domain.ProjectRepository {
	return &ProjectRepository{}
}

type ProjectRepository struct {
}

func (r *ProjectRepository) List(tx domain.Transaction, userID uint) ([]*domain.Project, error) {
	db := tx.(*gorm.DB)
	var projects []model.Project
	err := db.Where("user_id = ?", userID).
		Order("id ASC").
		Find(&projects).
		Error
	if err != nil {
		return nil, err
	}

	domains := make([]*domain.Project, len(projects))
	for i, project := range projects {
		domains[i] = project.ToDomain()
	}

	return domains, nil
}

func (r *ProjectRepository) Create(tx domain.Transaction, project *domain.Project) (*domain.Project, error) {
	db := tx.(*gorm.DB)
	model := model.FromDomainProject(project)
	if err := db.Create(model).Error; err != nil {
		return nil, err
	}
	return model.ToDomain(), nil
}

func (r *ProjectRepository) Get(tx domain.Transaction, userID, id uint) (*domain.Project, error) {
	db := tx.(*gorm.DB)
	var project model.Project
	if err := db.First(&project, id).Error; err != nil {
		return nil, err
	}
	return project.ToDomain(), nil
}

func (r *ProjectRepository) Update(tx domain.Transaction, project *domain.Project) (*domain.Project, error) {
	db := tx.(*gorm.DB)
	model := model.FromDomainProject(project)
	if err := db.Save(model).Error; err != nil {
		return nil, err
	}
	return model.ToDomain(), nil
}

func (r *ProjectRepository) Delete(tx domain.Transaction, userID, id uint) error {
	db := tx.(*gorm.DB)
	return db.Delete(&model.Project{}, id).Error
}
