package subscriber

import (
	"fmt"
	"time"

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
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		err := s.projectUseCase.CreateDemoProject(event.UserID)
		if err == nil {
			return
		}

		if i == maxRetries-1 {
			fmt.Printf("Failed to create demo project after %d retries for user %d: %v\n", maxRetries, event.UserID, err)
			return
		}

		// Exponential backoff (1秒、2秒、4秒)
		backoffDuration := time.Duration(1<<i) * time.Second
		time.Sleep(backoffDuration)
	}
}
