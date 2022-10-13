package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	/* ファイバーインスタンスを作成する
	ポート5000でHTTPリクエストをリッスンする。*/
	app := fiber.New()
	app.Use(cors.New())

	/* エンドポイントにアクセスするためのベースURLがhttp://localhost:5000/apiになる。 */
	api := app.Group("/api")

	// Test handler
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("App running 2")
	})

	log.Fatal(app.Listen(":5000"))
}
