package main

import (
	"log"
	"net/http"

	"burny-api/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Echoのインスタンス作成
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// PostgreSQLの接続情報
	dsn := "host=localhost user=yourusername password=yourpassword dbname=yourdbname port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal(err)
	}

	// マイグレーション
	if err := db.AutoMigrate(
		&models.Project{},
		&models.Sprint{},
		&models.SprintStat{},
	); err != nil {
		log.Fatalf("could not migrate: %v", err)
	}

	// ルートエンドポイント
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// エンドポイントの追加...

	// サーバーの開始
	e.Logger.Fatal(e.Start(":1323"))
}
