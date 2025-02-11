package infrastructure

import (
	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/model"
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

func (r *ProjectRepository) Create(project *domain.Project) (*domain.Project, error) {
	model := model.FromDomainProject(project)
	if err := DB.Create(model).Error; err != nil {
		return nil, err
	}

	return r.Get(model.ID)
}

func (r *ProjectRepository) Get(id uint) (*domain.Project, error) {
	var project model.Project
	if err := DB.First(&project, id).Error; err != nil {
		return nil, err
	}

	return project.ToDomain(), nil
}

func (r *ProjectRepository) Update(project *domain.Project) (*domain.Project, error) {
	model := model.FromDomainProject(project)
	if err := DB.Save(model).Error; err != nil {
		return nil, err
	}

	return r.Get(model.ID)
}

func (r *ProjectRepository) Delete(id uint) error {
	return DB.Delete(&model.Project{}, id).Error
}
