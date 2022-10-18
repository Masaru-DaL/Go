package router

import (
	"os"

	"github.com/labstack/echo/v4/middleware"

	_ "net/http"

	"github.com/labstack/echo/v4"
)

/* Routingを設定する関数 */
// 引数はecho.echo型で、戻り値はerror型
func SetRouter(e *echo.Echo) error {
	// APIが叩かれた時にログを出す。
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${host} ${method} ${uri} ${status} ${header}\n",
		Output: os.Stdout,
	}))
	// 予想外のエラーが発生した際にもサーバを落とさないようにする
	e.Use(middleware.Recover())
	// CORSに対応する
	e.Use(middleware.CORS())

	/* APIを書く */
	e.GET("/api/tasks", GetTaskHandler)

	/* port8000を開く */
	err := e.Start(":8000")
	return err
}
