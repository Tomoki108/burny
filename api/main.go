package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Tomoki108/burny/config"
	"github.com/Tomoki108/burny/di"
	"github.com/Tomoki108/burny/docs"
	"github.com/Tomoki108/burny/handler"
	"github.com/Tomoki108/burny/infrastructure"
	"github.com/golang-jwt/jwt/v5"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           Burny API
// @version         1.0
// @description     API Doc of Burny Backend
// @license.name  AGPL 3.0
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

	// DIコンテナからハンドラーを取得
	di.ProvideDependencies()
	var authH handler.AuthHandler
	di.Container.Invoke(func(h handler.AuthHandler) {
		authH = h
	})
	var projectH handler.ProjectHandler
	di.Container.Invoke(func(h handler.ProjectHandler) {
		projectH = h
	})
	var sprintH handler.SprintHandler
	di.Container.Invoke(func(h handler.SprintHandler) {
		sprintH = h
	})

	// ルーティング
	g := e.Group("/api/v1")
	g.POST("/sign_up", authH.SignUp)
	g.POST("/sign_in", authH.SignIn)

	ug := g.Group("", customJWTMiddleware([]byte(config.Conf.JwtSecret)))
	ug.GET("/projects", projectH.List)
	ug.POST("/projects", projectH.Create)
	ug.GET("/projects/:project_id", projectH.Get)
	ug.PUT("/projects/:project_id", projectH.Update)
	ug.DELETE("/projects/:project_id", projectH.Delete)
	ug.GET("/projects/:project_id/sprints", sprintH.List)
	ug.PUT("/sprints/:project_id/sprints/:sprint_id", sprintH.Update)

	// サーバー起動
	e.Logger.Fatal(e.Start(":1323"))
}

func customJWTMiddleware(secretKey []byte) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: secretKey,
		SuccessHandler: func(c echo.Context) {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			if userID, ok := claims["user_id"].(string); ok {
				userIDInt, _ := strconv.Atoi(userID)
				c.Set("user_id", uint(userIDInt))
			} else if userIDFloat, ok := claims["user_id"].(float64); ok {
				c.Set("user_id", uint(userIDFloat)) // 数値の場合
			}
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
		},
	})
}
