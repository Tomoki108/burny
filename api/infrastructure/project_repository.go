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

func (r *ProjectRepository) List() ([]*domain.Project, error) {
	var projects []model.Project
	if err := DB.Find(&projects).Error; err != nil {
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

	return r.Get(tx, model.ID)
}

func (r *ProjectRepository) Get(tx domain.Transaction, id uint) (*domain.Project, error) {
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

	return r.Get(tx, model.ID)
}

func (r *ProjectRepository) Delete(tx domain.Transaction, id uint) error {
	db := tx.(*gorm.DB)
	return db.Delete(&model.Project{}, id).Error
}
