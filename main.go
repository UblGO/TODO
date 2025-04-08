package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var Cfg *Config

func main() {
	Cfg = LoadConfig()
	//создание БД при необходимости
	if Cfg.CreateDB == "true" {
		createDataBase()
	}
	createTable()

	app := fiber.New()

	app.Get("/tasks", getTasksHandler)
	app.Post("/tasks", postTaskhandler)
	app.Put("/tasks/:id", updateTaskHandler)
	app.Delete("/tasks/:id", deleteTaskhandler)
	app.Listen(":" + Cfg.Port)
}
