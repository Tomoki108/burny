package usecase

import (
	"github.com/Tomoki108/burny/domain"
)

type ProjectUseCase struct {
	ProjectRepo   domain.ProjectRepository
	SprintRepo    domain.SprintRepository
	Transactioner domain.Transactioner
}

func (u ProjectUseCase) List() ([]*domain.Project, error) {
	return u.ProjectRepo.List()
}

func (u ProjectUseCase) Create(project *domain.Project) (*domain.Project, error) {
	var createdProject *domain.Project
	err := u.Transactioner.Transaction(func(tx domain.Transaction) (err error) {
		createdProject, err = u.ProjectRepo.Create(tx, project)
		if err != nil {
			return err
		}

		sprints := make([]*domain.Sprint, 0, createdProject.SprintCount)
		idealSP := createdProject.TotalSP / createdProject.SprintCount
		startDate := createdProject.StartDate
		endDate := startDate.AddDate(0, 0, 7*createdProject.SprintDuration)
		for i := 0; i < createdProject.SprintCount; i++ {
			sprint := &domain.Sprint{
				ProjectID: createdProject.ID,
				IdealSP:   idealSP,
				StartDate: startDate,
				EndDate:   endDate,
			}
			sprints = append(sprints, sprint)

			startDate = endDate
			endDate = startDate.AddDate(0, 0, 7*createdProject.SprintDuration)
		}

		for _, sprint := range sprints {
			_, err := u.SprintRepo.Create(tx, sprint)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return createdProject, err
}

func (u ProjectUseCase) Get(id uint) (*domain.Project, error) {
	return u.ProjectRepo.Get(u.Transactioner.Default(), id)
}

func (u ProjectUseCase) Update(project *domain.Project) (*domain.Project, error) {
	return u.ProjectRepo.Update(u.Transactioner.Default(), project)
}

func (u ProjectUseCase) Delete(id uint) error {
	return u.ProjectRepo.Delete(u.Transactioner.Default(), id)
}
