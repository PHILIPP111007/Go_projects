package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var db *pgx.Conn

func initDB() {
	var err error
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	db, err = pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
}

func createTask(c *fiber.Ctx) error {
	task := new(Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	query := "INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id"
	err := db.QueryRow(context.Background(), query, task.Title, task.Description, task.Status).Scan(&task.ID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create task"})
	}

	return c.Status(http.StatusCreated).JSON(task)
}

func getTasks(c *fiber.Ctx) error {
	rows, err := db.Query(context.Background(), "SELECT id, title, description, status, created_at, updated_at FROM tasks")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch tasks"})
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to scan task"})
		}
		tasks = append(tasks, task)
	}

	return c.JSON(tasks)
}

func updateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	task := new(Task)

	if err := c.BodyParser(task); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	query := "UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = now() WHERE id = $4"
	_, err := db.Exec(context.Background(), query, task.Title, task.Description, task.Status, id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update task"})
	}

	return c.JSON(task)
}

func deleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	query := "DELETE FROM tasks WHERE id = $1"
	_, err := db.Exec(context.Background(), query, id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete task"})
	}

	return c.SendStatus(http.StatusNoContent)
}

func main() {
	initDB()

	app := fiber.New()

	app.Get("/tasks", getTasks)
	app.Put("/tasks/:id", updateTask)
	app.Post("/tasks", createTask)
	app.Delete("/tasks/:id", deleteTask)

	log.Fatal(app.Listen(":3000"))
}
