package server

import (
	"github.com/Tomoki108/burny/config"
	"github.com/Tomoki108/burny/docs"
	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/handler"
	"github.com/Tomoki108/burny/middleware"
	"github.com/Tomoki108/burny/subscriber"
	"github.com/asaskevich/EventBus"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewEchoServer() *echo.Echo {
	// Echo インスタンス生成、全体に適用するミドルウェア設定
	e := echo.New()
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORS()) // デフォルトのCORS設定。（これがないとlocalhostの別のポートからの通信が許可されない）

	// DIコンテナからイベントサブスクライバーを取得
	var userEventSub subscriber.UserEventSubscriber
	Container.Invoke(func(s subscriber.UserEventSubscriber) {
		userEventSub = s
	})

	// API DOC
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	docs.SwaggerInfo.Host = config.Conf.Host

	// イベントサブスクライバー登録
	var bus EventBus.Bus
	Container.Invoke(func(b EventBus.Bus) {
		bus = b
	})
	bus.Subscribe(domain.UserCreatedTopic, userEventSub.HandleUserCreatedEvent)

	// DIコンテナからハンドラーを取得
	var authH handler.AuthHandler
	var projectH handler.ProjectHandler
	var sprintH handler.SprintHandler
	var apiKeyH handler.APIKeyHandler
	Container.Invoke(func(
		ah handler.AuthHandler,
		ph handler.ProjectHandler,
		sh handler.SprintHandler,
		aph handler.APIKeyHandler,
	) {
		authH = ah
		projectH = ph
		sprintH = sh
		apiKeyH = aph
	})

	// 認証ミドルウェアをDIコンテナから取得
	var apiKeyAuth *middleware.APIKeyAuthMiddleware
	var jwtAuth *middleware.JWTAuthMiddleware
	Container.Invoke(func(a *middleware.APIKeyAuthMiddleware, j *middleware.JWTAuthMiddleware) {
		apiKeyAuth = a
		jwtAuth = j
	})

	// ルーティング
	g := e.Group("/api/v1")
	g.POST("/sign_up", authH.SignUp)
	g.POST("/sign_in", authH.SignIn)

	ug := g.Group("", jwtAuth.Middleware(), apiKeyAuth.Middleware())
	ug.GET("/projects", projectH.List)
	ug.POST("/projects", projectH.Create)
	ug.GET("/projects/:project_id", projectH.Get)
	ug.PUT("/projects/:project_id", projectH.Update)
	ug.DELETE("/projects/:project_id", projectH.Delete)
	ug.GET("/projects/:project_id/sprints", sprintH.List)
	ug.PATCH("/projects/:project_id/sprints/:sprint_id", sprintH.Update)
	ug.GET("/apikeys/status", apiKeyH.CheckStatus)
	ug.POST("/apikeys", apiKeyH.Create)
	ug.DELETE("/apikeys", apiKeyH.Delete)

	return e
}
