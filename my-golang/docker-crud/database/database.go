package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name string `json:"name"`
	Task string `json:"task"`
}

func InitDatabase() error {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&Todo{})

	return nil
}

func GetAllTasks() ([]Todo, error) {
	var todos []Todo

	db, err := gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{})
	if err != nil {
		return todos, err
	}

	db.Find(&todos)

	return todos, nil
}

func CreateTodo(name string, task string) (Todo, error) {
	var newTodo = Todo{Name: name, Task: task}

	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		return newTodo, err
	}
	db.Create(&Todo{Name: name, Task: task})

	return newTodo, nil
}
