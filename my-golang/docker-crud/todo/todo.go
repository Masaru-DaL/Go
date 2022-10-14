package todo

import (
	"docker-crud/database"

	"github.com/gofiber/fiber/v2"
)

func GetAllTasks(c *fiber.Ctx) error {
	result, err := database.GetAllTasks()
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
}

func SaveTask(c *fiber.Ctx) error {
	newTodo := new(database.Todo)

	err := c.BodyParser(newTodo)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	result, err := database.CreateTodo(newTodo.Name, newTodo.Task)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	c.Status(200).JSON(&fiber.Map{
		"success": false,
		"message": "",
		"data":    result,
	})

	return nil
}
