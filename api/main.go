package main

import (
	"log"

	"github.com/Tomoki108/burny/config"
	"github.com/Tomoki108/burny/docs"
	"github.com/Tomoki108/burny/handler"
	"github.com/Tomoki108/burny/infrastructure"
	"github.com/Tomoki108/burny/usecase"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           Burny API
// @version         1.0
// @description     API Doc of Burny Backend
// @license.name  GPL 3.0
// @license.url   https://www.gnu.org/licenses/agpl-3.0.en.html
// @BasePath  /api/v1
func main() {
	// 環境変数の読み込み
	if err := config.Init(); err != nil {
		log.Fatal(err.Error())
	}

	// DB接続
	if err := infrastructure.ConnectDB(); err != nil {
		log.Fatal(err.Error())
	}

	// Echo インスタンス生成、全体に適用するミドルウェア設定
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS()) // デフォルトのCORS設定。（これがないとlocalhostの別のポートからの通信が許可されない）

	// API DOC
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	docs.SwaggerInfo.Host = config.Conf.Host

	// ルーティング
	projectHandler := handler.ProjectHandler{
		UseCase: usecase.ProjectUseCase{
			ProjectRepo:   infrastructure.NewProjectRepository(),
			SprintRepo:    infrastructure.NewSprintRepository(),
			Transactioner: infrastructure.NewTransactioner(),
		},
	}
	sprintHandler := handler.SprintHandler{
		UseCase: usecase.SprintUseCase{
			SprintRepo:    infrastructure.NewSprintRepository(),
			Transactioner: infrastructure.NewTransactioner(),
		},
	}
	authHandler := handler.AuthHandler{
		Usecase: usecase.AuthUseCase{
			Repo:          infrastructure.NewUserRepository(),
			Transactioner: infrastructure.NewTransactioner(),
		},
	}

	g := e.Group("/api/v1")
	g.POST("/sign_up", authHandler.SignUp)
	g.POST("/sign_in", authHandler.SignIn)

	ug := g.Group("", echojwt.JWT([]byte([]byte(config.Conf.JwtSecret))))
	ug.GET("/projects", projectHandler.List)
	ug.POST("/projects", projectHandler.Create)
	ug.GET("/projects/:project_id", projectHandler.Get)
	ug.PUT("/projects/:project_id", projectHandler.Update)
	ug.DELETE("/projects/:project_id", projectHandler.Delete)
	ug.GET("/projects/:project_id/sprints", sprintHandler.List)
	ug.PUT("/sprints/:project_id/sprints/:sprint_id", sprintHandler.Update)

	// サーバー起動
	e.Logger.Fatal(e.Start(":1323"))
}
