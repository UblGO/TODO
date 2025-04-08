package main

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Структура задачи для сериализации/десериализации
type Task struct {
	Id          int       `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status,omitempty"`
	Created_at  time.Time `json:"created_at,omitempty"`
	Updated_at  time.Time `json:"updated_at,omitempty"`
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

// Хендлер для получения всех задач
func getTasksHandler(c *fiber.Ctx) error {
	t, err := getTasks()
	if err != nil {
		return err
	}
	return c.JSON(t)
}

// Хендлер для создания задачи
func postTaskhandler(c *fiber.Ctx) error {
	var task Task
	err := c.BodyParser(&task)
	if err != nil {
		return err
	}
	err = postTask(task)
	if err != nil {
		return err
	}
	return nil
}

// Хендлер для обновления задачи
func updateTaskHandler(c *fiber.Ctx) error {
	idString := c.Params("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}
	if err = updateTask(id); err != nil {
		return err
	}
	return nil
}

// Хендлер для удаления задачи
func deleteTaskhandler(c *fiber.Ctx) error {
	idString := c.Params("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}
	if err = deleteTask(id); err != nil {
		return err
	}
	return nil
}
