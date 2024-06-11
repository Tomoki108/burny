package main

import (
	"log"
	"net/http"

	"burny-api/db"
	"burny-api/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echoのインスタンス作成
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// DB接続
	if err := db.Connect(); err != nil {
		log.Fatalf(err.Error())
	}

	// ルートエンドポイント
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// エンドポイントの追加...
	e.GET("/projects", handlers.ListProjectsHandler)
	e.POST("/projects", handlers.CreateProjectHandler)
	e.GET("/projects/:id", handlers.GetProjectHandler)
	e.PUT("/projects/:id", handlers.UpdateProjectHandler)
	e.DELETE("/projects/:id", handlers.DeleteProjectHandler)

	// サーバーの開始
	e.Logger.Fatal(e.Start(":1323"))
}
