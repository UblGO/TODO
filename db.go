package main

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
)

// Sql-запросы
var (
	getTasksQuery    = "SELECT * FROM scheduler ORDER BY id;"
	deleteTaskQuery  = "DELETE FROM scheduler WHERE id = $1;"
	postTaskQuery    = "INSERT INTO scheduler (title, description) VALUES ($1, $2);"
	createTableQuery = `CREATE TABLE scheduler (id SERIAL PRIMARY KEY, title TEXT NOT NULL, description TEXT,
						   status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new', created_at TIMESTAMP DEFAULT now(),
						   updated_at TIMESTAMP DEFAULT now());`
	updateTaskQuery = `UPDATE scheduler
	SET status = 
	  CASE
		WHEN status = 'new' THEN 'in_progress'
		when status = 'in_progress' THEN 'done'
		ELSE status
	  END,
	  updated_at=now() WHERE id = $1;`
)

// Проверяет наличие таблицы
func tableExists(conn *pgx.Conn, tableName string) (bool, error) {
	var n int64
	err := conn.QueryRow(context.Background(), "SELECT 1 FROM information_schema.tables WHERE table_name = $1", tableName).Scan(&n)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// Создает таблицу, если таблица уже существует ничего не делает
func createTable() {
	db, err := pgx.Connect(context.Background(), Cfg.DB.ConnInfo())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())
	exists, err := tableExists(db, "scheduler")
	if err != nil {
		log.Fatal(err)
	}
	if exists {
		return
	}
	_, err = db.Exec(context.Background(), createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

}

// Cоздание БД при необходимости
func createDataBase() {
	db, err := pgx.Connect(context.Background(), Cfg.DB.Pgconn())
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(context.Background(), "create database "+Cfg.DB.DbName)
	if err != nil {
		log.Fatal(err)
	}
	db.Close(context.Background())
}

// Добавляет задачу в БД
func postTask(task Task) error {
	db, err := pgx.Connect(context.Background(), Cfg.DB.ConnInfo())
	if err != nil {
		return err
	}
	defer db.Close(context.Background())

	res, err := db.Exec(context.Background(), postTaskQuery, task.Title, task.Description)
	if err != nil {
		return err
	}
	if res.RowsAffected() != 1 {
		return errors.New("Task was not posted")
	}
	return nil
}

// Получает список всех задач из БД
func getTasks() (Tasks, error) {
	db, err := pgx.Connect(context.Background(), Cfg.DB.ConnInfo())
	if err != nil {
		return Tasks{}, err
	}
	defer db.Close(context.Background())

	rows, err := db.Query(context.Background(), getTasksQuery)
	if err != nil {
		return Tasks{}, err
	}
	var tasks Tasks
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Status, &t.Created_at, &t.Updated_at)
		if err != nil {
			return tasks, err
		}
		tasks.Tasks = append(tasks.Tasks, t)
	}
	return tasks, nil
}

// Обновляет задачу в БД
func updateTask(id int) error {
	db, err := pgx.Connect(context.Background(), Cfg.DB.ConnInfo())
	if err != nil {
		return err
	}
	defer db.Close(context.Background())

	res, err := db.Exec(context.Background(), updateTaskQuery, id)
	if err != nil {
		return err
	}
	if res.RowsAffected() != 1 {
		return errors.New("Task not found")
	}
	return nil
}

// Удаляет задачу bp БД
func deleteTask(id int) error {
	db, err := pgx.Connect(context.Background(), Cfg.DB.ConnInfo())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	res, err := db.Exec(context.Background(), deleteTaskQuery, id)
	if err != nil {
		return err
	}
	if res.RowsAffected() != 1 {
		return errors.New("Task not found")
	}
	return nil
}
