package main

import (
	"log"
	"net/http"

	"github.com/Tomoki108/burny/db"
	"github.com/Tomoki108/burny/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type data struct {
	Date        string `json:"date"`
	SolarPanels int    `json:"SolarPanels"`
	Inverters   int    `json:"Inverters"`
}

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
	e.GET("test", func(c echo.Context) error {
		list := []data{
			{Date: "Jan 22", SolarPanels: 2890, Inverters: 2338},
			{Date: "Feb 22", SolarPanels: 2756, Inverters: 2103},
			{Date: "Mar 22", SolarPanels: 3322, Inverters: 2194},
			{Date: "Apr 22", SolarPanels: 3470, Inverters: 2108},
			{Date: "May 22", SolarPanels: 3475, Inverters: 1812},
			{Date: "Jun 22", SolarPanels: 3129, Inverters: 1726},
			{Date: "Jul 22", SolarPanels: 3490, Inverters: 1982},
			{Date: "Aug 22", SolarPanels: 2903, Inverters: 2012},
		}

		return c.JSON(http.StatusOK, list)
	})

	// CORSの設定。localhost（別のポート）からの通信を許可するだけならこれで十分らしい。
	// https://qiita.com/sola-msr/items/828e2eb45cf05b1a2ad4
	e.Use(middleware.CORS())
	// サーバーの開始
	e.Logger.Fatal(e.Start(":1323"))
}
