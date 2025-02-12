package main

import (
	"log"

	_ "github.com/Tomoki108/burny/docs"
	"github.com/Tomoki108/burny/handler"
	"github.com/Tomoki108/burny/infrastructure"
	"github.com/Tomoki108/burny/usecase"

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
	e := echo.New()

	// ミドルウェア
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS()) // デフォルトのCORS設定。（これがないとlocalhost（別のポート）からの通信が許可されない）

	// DB接続
	if err := infrastructure.ConnectDB(); err != nil {
		log.Fatal(err.Error())
	}

	// ルーティング
	projectHandler := handler.ProjectHandler{
		Repo: infrastructure.NewProjectRepository(),
	}
	sprintHandler := handler.SprintHandler{
		Repo: infrastructure.NewSprintRepository(),
	}

	g := e.Group("/api/v1")
	g.GET("/swagger/*", echoSwagger.WrapHandler)

	authHandler := handler.AuthHandler{
		Usecase: usecase.SignUpUseCase{
			Repo: infrastructure.NewUserRepository(),
		},
	}
	g.POST("/sign_up", authHandler.SignUp)
	g.POST("/sign_in", authHandler.SignIn)

	g.GET("/projects", projectHandler.List)
	g.POST("/projects", projectHandler.Create)
	g.GET("/projects/:id", projectHandler.Get)
	g.PUT("/projects/:id", projectHandler.Update)
	g.DELETE("/projects/:id", projectHandler.Delete)
	g.GET("/sprints", sprintHandler.List)
	g.PUT("/sprints/:id", sprintHandler.Update)

	// サーバーの開始
	e.Logger.Fatal(e.Start(":1323"))
}
