package main

import (
	"docker-crud/database"
	"docker-crud/todo"
	"log"

	"github.com/gofiber/fiber/v2"
)

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running!! Send your request")
}

func setupRoutes(app *fiber.App) {

	app.Get("/", status)

	app.Get("/api/todo", todo.GetAllTasks)
	app.Post("/api/todo", todo.SaveTask)
}

func main() {
	app := fiber.New()
	dbErr := database.InitDatabase()

	if dbErr != nil {
		panic(dbErr)
	}

	setupRoutes(app)

	log.Fatal(app.Listen(":1323"))
}
