package usecase

import (
	"errors"
	"time"

	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/handler/io"
)

var ErrProjectNotFound = errors.New("project not found")
var ErrSprintHasAlreadyStarted = errors.New("can not delete sprint that has already started")

type ProjectUseCase struct {
	ProjectRepo   domain.ProjectRepository
	SprintRepo    domain.SprintRepository
	Transactioner domain.Transactioner
}

func NewProjectUseCase(
	projectRepo domain.ProjectRepository,
	sprintRepo domain.SprintRepository,
	transactioner domain.Transactioner,
) ProjectUseCase {
	return ProjectUseCase{
		ProjectRepo:   projectRepo,
		SprintRepo:    sprintRepo,
		Transactioner: transactioner,
	}
}

func (u ProjectUseCase) List(userID uint) ([]*domain.Project, error) {
	return u.ProjectRepo.List(u.Transactioner.Default(), userID)
}

func (u ProjectUseCase) Create(userID uint, req io.CreateProjectRequest) (*domain.Project, error) {
	project := &domain.Project{
		UserID:         userID,
		Title:          req.Title,
		Description:    req.Description,
		TotalSP:        req.TotalSP,
		StartDate:      req.StartDate,
		SprintDuration: req.SprintDuration,
		SprintCount:    req.SprintCount,
	}

	err := u.Transactioner.Transaction(func(tx domain.Transaction) (err error) {
		project, err = u.ProjectRepo.Create(tx, project)
		if err != nil {
			return err
		}

		sprints := make([]*domain.Sprint, 0, project.SprintCount)
		idealSP := project.TotalSP / project.SprintCount
		startDate := project.StartDate
		endDate := startDate.AddDate(0, 0, 7*project.SprintDuration-1)
		for i := 0; i < project.SprintCount; i++ {
			sprint := &domain.Sprint{
				UserID:    userID,
				ProjectID: project.ID,
				IdealSP:   idealSP,
				StartDate: startDate,
				EndDate:   endDate,
			}
			sprints = append(sprints, sprint)

			startDate = endDate.AddDate(0, 0, 1)
			endDate = startDate.AddDate(0, 0, 7*project.SprintDuration-1)
		}

		for _, sprint := range sprints {
			err := u.SprintRepo.Create(tx, sprint)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return project, err
}

func (u ProjectUseCase) Get(userID, id uint) (*domain.Project, error) {
	return u.ProjectRepo.Get(u.Transactioner.Default(), userID, id)
}

func (u ProjectUseCase) Update(userID uint, req io.UpdateProjectRequest) (*domain.Project, error) {
	var updatedProject *domain.Project
	err := u.Transactioner.Transaction(func(tx domain.Transaction) (err error) {
		project, err := u.ProjectRepo.Get(tx, userID, req.ProjectID)
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
			endDate := startDate.AddDate(0, 0, 7*project.SprintDuration-1)
			for i := 0; i < countDiff; i++ {
				sprint := &domain.Sprint{
					UserID:    project.UserID,
					ProjectID: project.ID,
					StartDate: startDate,
					EndDate:   endDate,
				}
				sprints = append(sprints, sprint)

				startDate = endDate.AddDate(0, 0, 1)
				endDate = startDate.AddDate(0, 0, 7*project.SprintDuration-1)
			}
		} else if countDiff < 0 {
			sprints = sprints[:req.SprintCount]
			if sprints[len(sprints)-1].EndDate.Before(time.Now()) {
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
		}
		_, err = u.SprintRepo.UpsertList(tx, sprints)
		if err != nil {
			return err
		}

		updatedProject, err = u.ProjectRepo.Update(tx, project)
		return err
	})

	return updatedProject, err
}

func (u ProjectUseCase) Delete(userID uint, req io.DeleteProjectRequest) error {
	return u.Transactioner.Transaction(func(tx domain.Transaction) (err error) {
		sprints, err := u.SprintRepo.List(tx, req.ProjectID)
		if err != nil {
			return err
		}
		for _, sprint := range sprints {
			u.SprintRepo.Delete(tx, sprint.ProjectID, sprint.ID)
			if err != nil {
				return
			}
		}
		return u.ProjectRepo.Delete(u.Transactioner.Default(), userID, req.ProjectID)
	})
}

func (u ProjectUseCase) CreateDemoProject(userID uint) error {
	// startDateに4週間前の月曜日をセットする
	now := time.Now()
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	startDate = startDate.AddDate(0, 0, -28)
	for startDate.Weekday() != time.Monday {
		startDate = startDate.AddDate(0, 0, -1)
	}

	project := &domain.Project{
		UserID:         userID,
		Title:          "Demo Project",
		Description:    "This is a demo project to show how burny works. Check this out!!",
		TotalSP:        120,
		StartDate:      startDate,
		SprintDuration: 1,
		SprintCount:    6,
	}
	actualSPs := []int{15, 27, 10, 24, 0, 0}

	err := u.Transactioner.Transaction(func(tx domain.Transaction) (err error) {
		project, err = u.ProjectRepo.Create(tx, project)
		if err != nil {
			return err
		}

		sprints := make([]*domain.Sprint, 0, project.SprintCount)
		idealSP := project.TotalSP / project.SprintCount
		startDate := project.StartDate
		endDate := startDate.AddDate(0, 0, 7*project.SprintDuration-1)
		for i := 0; i < project.SprintCount; i++ {
			sprint := &domain.Sprint{
				UserID:    userID,
				ProjectID: project.ID,
				IdealSP:   idealSP,
				ActualSP:  actualSPs[i],
				StartDate: startDate,
				EndDate:   endDate,
			}
			sprints = append(sprints, sprint)

			startDate = endDate.AddDate(0, 0, 1)
			endDate = startDate.AddDate(0, 0, 7*project.SprintDuration-1)
		}

		for _, sprint := range sprints {
			err := u.SprintRepo.Create(tx, sprint)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
