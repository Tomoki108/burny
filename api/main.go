package main

import (
	"log"

	"github.com/Tomoki108/burny/config"
	"github.com/Tomoki108/burny/infrastructure"
	"github.com/Tomoki108/burny/server"
)

// @title                      Burny API
// @version                    1.0
// @description                Burny Backend API Doc. \nNOTE: JWT Auhtentication is also supported but not displayed due to swagger generaton tool limitation.
// @license.name               AGPL 3.0
// @license.url                https://www.gnu.org/licenses/agpl-3.0.en.html
// @securityDefinitions.apikey Account API Key
// @in                         Header
// @name                       X-API-Key
// @BasePath                   /api/v1
func main() {
	// 環境変数の読み込み
	if err := config.Init(); err != nil {
		log.Fatal(err.Error())
	}
	// DB接続
	if err := infrastructure.ConnectDB(); err != nil {
		log.Fatal(err.Error())
	}
	// AWS SES接続
	if err := infrastructure.ConnectAWSSES(); err != nil {
		log.Fatal(err.Error())
	}
	// DIコンテナの初期化
	server.InitDIContainer()
	// サーバーの起動
	e := server.NewEchoServer()
	e.Logger.Fatal(e.Start(":" + config.Conf.Port))
}
