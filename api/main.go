package main

import (
	"log"

	_ "github.com/Tomoki108/burny/docs"
	"github.com/Tomoki108/burny/handler"
	"github.com/Tomoki108/burny/infrastructure"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           Swagger Example API2
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Echoのインスタンス作成
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// DB接続
	if err := infrastructure.ConnectDB(); err != nil {
		log.Fatal(err.Error())
	}

	projectHandler := handler.ProjectHandler{
		Repo: infrastructure.NewProjectRepository(),
	}
	sprintHandler := handler.SprintHandler{
		Repo: infrastructure.NewSprintRepository(),
	}

	// API DOC
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// エンドポイントの追加
	e.GET("/projects", projectHandler.List)
	e.POST("/projects", projectHandler.Create)
	e.GET("/projects/:id", projectHandler.Get)
	e.PUT("/projects/:id", projectHandler.Update)
	e.DELETE("/projects/:id", projectHandler.Delete)

	e.GET("/sprints", sprintHandler.List)
	e.PUT("/sprints/:id", sprintHandler.Update)

	// CORSの設定。localhost（別のポート）からの通信を許可するだけならこれで十分らしい。
	// https://qiita.com/sola-msr/items/828e2eb45cf05b1a2ad4
	e.Use(middleware.CORS())
	// サーバーの開始
	e.Logger.Fatal(e.Start(":1323"))
}
