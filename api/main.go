package main

import (
	"log"

	"github.com/Tomoki108/burny/config"
	_ "github.com/Tomoki108/burny/docs"
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
// @host      temp.com
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

	// ルーティング
	projectHandler := handler.ProjectHandler{
		Usecase: usecase.ProjectUseCase{
			ProjectRepo: infrastructure.NewProjectRepository(),
		},
	}
	sprintHandler := handler.SprintHandler{
		Repo: infrastructure.NewSprintRepository(),
	}
	authHandler := handler.AuthHandler{
		Usecase: usecase.AuthUseCase{
			Repo: infrastructure.NewUserRepository(),
		},
	}

	g := e.Group("/api/v1")
	g.POST("/sign_up", authHandler.SignUp)
	g.POST("/sign_in", authHandler.SignIn)

	ug := e.Group("", echojwt.JWT([]byte([]byte(config.Conf.JwtSecret))))
	ug.GET("/projects", projectHandler.List)
	ug.POST("/projects", projectHandler.Create)
	ug.GET("/projects/:id", projectHandler.Get)
	ug.PUT("/projects/:id", projectHandler.Update)
	ug.DELETE("/projects/:id", projectHandler.Delete)
	ug.GET("/sprints", sprintHandler.List)
	ug.PUT("/sprints/:id", sprintHandler.Update)

	// サーバー起動
	e.Logger.Fatal(e.Start(":1323"))
}
