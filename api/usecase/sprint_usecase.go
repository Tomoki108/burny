package usecase

import (
	"errors"

	"github.com/Tomoki108/burny/domain"
)

var ErrSprintNotFound = errors.New("sprint not found")

type SprintUseCase struct {
	SprintRepo    domain.SprintRepository
	Transactioner domain.Transactioner
}

func (u SprintUseCase) List(pojectID uint) ([]*domain.Sprint, error) {
	return u.SprintRepo.List(u.Transactioner.Default(), pojectID)
}

func (u SprintUseCase) Update(projectID, sprintID uint, actualSP int) (*domain.Sprint, error) {
	return u.SprintRepo.Update(u.Transactioner.Default(), projectID, sprintID, actualSP)
}
