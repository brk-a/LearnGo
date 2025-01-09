package routes

import (
	"fmt"
	"log"
	"os"

	"github.com/brk-a/task_manager/controllers"
	"github.com/gofiber/fiber/v2"
)

// Routes sets up the HTTP routes for the task manager application.
// It initialises a new Fiber application, defines the endpoints for CRUD operations on tasks
// and starts the server listening on the specified port.
// If the PORT environment variable is not set, it defaults to 8080.
func Routes() {
	app := fiber.New()

	// GET /api/v1/tasks - Retrieve all tasks
	app.Get("/api/v1/tasks", controllers.GetTasks)

	// GET /api/v1/tasks/:id - Retrieve a specific task by ID
	app.Get("/api/v1/tasks/:id", controllers.GetTask)

	// POST /api/v1/tasks - Create a new task
	app.Post("/api/v1/tasks", controllers.CreateTask)

	// PATCH /api/v1/tasks/:id - Update an existing task by ID
	app.Patch("/api/v1/tasks/:id", controllers.UpdateTask)

	// DELETE /api/v1/tasks/:id - Delete a specific task by ID
	app.Delete("/api/v1/tasks/:id", controllers.DeleteTask)

	// Determine the listening port
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	// Start the server and log any fatal errors
	log.Fatal(app.Listen(fmt.Sprintf("0.0.0.0:%s", PORT)))
}
