package server

import (
	"github.com/Tomoki108/burny/config"
	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/handler"
	"github.com/Tomoki108/burny/infrastructure"
	"github.com/Tomoki108/burny/middleware"
	"github.com/Tomoki108/burny/subscriber"
	"github.com/Tomoki108/burny/usecase"
	"github.com/asaskevich/EventBus"
	"go.uber.org/dig"
)

var Container *dig.Container

type provideArg struct {
	constructor interface{}
	opts        []dig.ProvideOption
}

func InitDIContainer() {
	Container = dig.New()

	args := []provideArg{
		// handler
		{handler.NewAuthHandler, nil},
		{handler.NewProjectHandler, nil},
		{handler.NewSprintHandler, nil},
		{handler.NewAPIKeyHandler, nil},
		// middleware
		{middleware.NewAPIKeyAuthMiddleware, nil},
		{
			func() *middleware.JWTAuthMiddleware {
				return middleware.NewJWTAuthMiddleware([]byte(config.Conf.JwtSecret))
			},
			nil,
		},
		// subscriber
		{subscriber.NewUserEventSubscriber, nil},
		// usecase
		{usecase.NewProjectUseCase, nil},
		{usecase.NewSprintUseCase, nil},
		{usecase.NewAuthUseCase, nil},
		{usecase.NewAPIKeyUseCase, nil},
		// infra
		{infrastructure.NewUserRepository, []dig.ProvideOption{dig.As(new(domain.UserRepository))}},
		{infrastructure.NewProjectRepository, []dig.ProvideOption{dig.As(new(domain.ProjectRepository))}},
		{infrastructure.NewSprintRepository, []dig.ProvideOption{dig.As(new(domain.SprintRepository))}},
		{infrastructure.NewAPIKeyRepository, []dig.ProvideOption{dig.As(new(domain.APIKeyRepository))}},
		{infrastructure.NewTransactioner, []dig.ProvideOption{dig.As(new(domain.Transactioner))}},
		// other
		{EventBus.New, nil},
	}

	for _, arg := range args {
		if err := Container.Provide(arg.constructor, arg.opts...); err != nil {
			panic(err)
		}
	}
}
