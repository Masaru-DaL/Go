package todos

import "github.com/gofiber/fiber/v2"

func GetAllTasks(c *fiber.Ctx) error {
	return c.SendString("All Tasks")
}

func SaveTask(c *fiber.Ctx) error {
	return c.SendString("Task Saved!!")
}
