package main

import (
	"dokcer-crud/todos"
	"log"

	"github.com/gofiber/fiber/v2"
)

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running!! Send your request")
}

func setupRoutes(app *fiber.App) {

	app.Get("/", status)

	app.Get("/api/todos", todos.GetAllTasks)
	app.Post("/api/todos", todos.SaveTask)
}

func main() {
	app := fiber.New()
	setupRoutes(app)

	log.Fatal(app.Listen(":1323"))
}
