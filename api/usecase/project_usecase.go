package usecase

import (
	"errors"
	"time"

	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/handler/io"
)

var ErrSprintHasAlreadyStarted = errors.New("既に開始済みのスプリントが削除される様な更新はできません")

type ProjectUseCase struct {
	ProjectRepo   domain.ProjectRepository
	SprintRepo    domain.SprintRepository
	Transactioner domain.Transactioner
}

func (u ProjectUseCase) List(userID uint) ([]*domain.Project, error) {
	return u.ProjectRepo.List(u.Transactioner.Default(), userID)
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

func (u ProjectUseCase) Update(req io.UpdateProjectRequest) (*domain.Project, error) {
	var updatedProject *domain.Project
	err := u.Transactioner.Transaction(func(tx domain.Transaction) (err error) {
		project, err := u.ProjectRepo.Get(tx, req.PrrojectID)
		if err != nil {
			return err
		}
		sprints, err := u.SprintRepo.List(tx, project.ID)
		if err != nil {
			return err
		}

		countDiff := req.SprintCount - project.SprintCount
		project.Title = req.Title
		project.Description = req.Description
		project.SprintCount = req.SprintCount
		project.TotalSP = req.TotalSP

		if countDiff > 0 {
			startDate := sprints[len(sprints)-1].EndDate
			endDate := startDate.AddDate(0, 0, 7*project.SprintDuration)
			for i := 0; i < countDiff; i++ {
				sprint := &domain.Sprint{
					ProjectID: project.ID,
					StartDate: startDate,
					EndDate:   endDate,
				}
				sprints = append(sprints, sprint)

				startDate = endDate
				endDate = startDate.AddDate(0, 0, 7*project.SprintDuration)
			}
		} else if countDiff < 0 {
			sprints = sprints[:req.SprintCount]
			if sprints[len(sprints)-1].EndDate.After(time.Now()) {
				return ErrSprintHasAlreadyStarted
			}

			for _, sprint := range sprints[req.SprintCount:] {
				err := u.SprintRepo.Delete(tx, sprint.ProjectID, sprint.ID)
				if err != nil {
					return err
				}
			}
		}

		idealSP := project.TotalSP / req.SprintCount
		for _, sprint := range sprints {
			sprint.IdealSP = idealSP
			_, err := u.SprintRepo.Update(tx, sprint.ProjectID, sprint.ID, sprint.IdealSP)
			if err != nil {
				return err
			}
		}

		updatedProject, err = u.ProjectRepo.Update(tx, project)
		return err
	})

	return updatedProject, err
}

func (u ProjectUseCase) Delete(id uint) error {
	return u.ProjectRepo.Delete(u.Transactioner.Default(), id)
}
