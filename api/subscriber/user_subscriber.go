package subscriber

import (
	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/usecase"
)

type UserEventSubscriber struct {
	projectUseCase usecase.ProjectUseCase
}

func NewUserEventSubscriber(projectUseCase usecase.ProjectUseCase) UserEventSubscriber {
	return UserEventSubscriber{
		projectUseCase: projectUseCase,
	}
}

func (s *UserEventSubscriber) HandleUserCreatedEvent(event domain.UserCreatedEvent) {
	s.projectUseCase.CreateDemoProject(event.UserID)
}
