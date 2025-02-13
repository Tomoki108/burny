package usecase

import "github.com/Tomoki108/burny/domain"

type ProjectUseCase struct {
	Repo domain.ProjectRepository
}

func (u ProjectUseCase) List() ([]*domain.Project, error) {
	return u.Repo.List()
}

func (u ProjectUseCase) Create(project *domain.Project) (*domain.Project, error) {
	return u.Repo.Create(project)
}

func (u ProjectUseCase) Get(id uint) (*domain.Project, error) {
	return u.Repo.Get(id)
}

func (u ProjectUseCase) Update(project *domain.Project) (*domain.Project, error) {
	return u.Repo.Update(project)
}

func (u ProjectUseCase) Delete(id uint) error {
	return u.Repo.Delete(id)
}
