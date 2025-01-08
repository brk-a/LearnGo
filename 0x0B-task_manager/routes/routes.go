package routes

import (
	"fmt"
	"log"
	"os"

	"github.com/brk-a/task_manager/controllers"
	"github.com/gofiber/fiber"
)

func Routes() {
	app := fiber.New()
	app.Get("/api/v1/tasks", controllers.GetTasks)
	app.Get("/api/v1/tasks/:id", controllers.GetTask)
	app.Post("/api/v1/tasks", controllers.CreateTask)
	app.Patch("/api/v1/tasks/:id", controllers.UpdateTask)
	app.Delete("/api/v1/tasks/:id", controllers.DeleteTask)

	PORT := os.Getenv("PORT")
	if PORT=="" {
		PORT = "8080"
	}
	log.Fatal(app.Listen(fmt.Sprintf("0.0.0.0:%s", PORT)))
}
