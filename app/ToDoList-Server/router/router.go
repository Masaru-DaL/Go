package router

import "os"

/* Routingを設定する関数 */
// 引数はecho.echo型で、戻り値はerror型
func SetRouter(e *echo.Echo) error {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${host} ${method} ${uri} ${status} ${header}\n",
		Output: os.Stdout,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	/* APIを書く */

	/* port8000を開く */
	err := e.Start(":8000")
	return err
}
