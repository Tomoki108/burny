package usecase

import "github.com/Tomoki108/burny/domain"

type SprintUseCase struct {
	SprintRepo    domain.SprintRepository
	Transactioner domain.Transactioner
}

func (u SprintUseCase) List() ([]*domain.Sprint, error) {
	return u.SprintRepo.List(u.Transactioner.New())
}

func (u SprintUseCase) Update(sprint *domain.Sprint) (*domain.Sprint, error) {
	return u.SprintRepo.Update(u.Transactioner.New(), sprint)
}
