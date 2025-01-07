package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type User struct {}

type Todo struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Completed bool `json:"completed"`
	Duration int64 `json:"duration"`
	// Created_at int64 `json:"created_at"`
	// Updated_at int64 `json:"updated_at"`
	// Deleted_at int64 `json:"deleted_at"`
	// User User `json:"user"`
	Due_date int64 `json:"due_date"`
}

func main(){
	app := fiber.New()

	todos := []Todo{}

	// routes
	app.Get("/api/tasks", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"msg": "OK",
			"data": todos,
		})
	})
	app.Get("/api/tasks:id", func(c *fiber.Ctx) error {
		todoId := c.Params("id")
		for _, todo  := range(todos){
			if fmt.Sprint(todo.ID) == todoId {
                return c.Status(200).JSON(fiber.Map{
                    "msg": "OK",
                    "data": todo,
                })
            }
		}
		return c.Status(500).JSON(fiber.Map{"error": "could not retrieve"})
	})
	app.Post("/api/tasks", func (c *fiber.Ctx) error {
		todo := &Todo{}
		if err := c.BodyParser(todo); err!= nil {
            return c.Status(500).JSON(fiber.Map{"error": err.Error()})
        }
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Body must be a string"})
		}

		todo.ID = int64(len(todos)) + 1
		todos = append(todos, *todo)
		
		return c.Status(201).JSON(fiber.Map{"msg": "created successfully"})
		
	})
	app.Patch("/api/tasks:id", func (c *fiber.Ctx) error  {
		todoId := c.Params("id")
		for i, todo  := range(todos){
			if fmt.Sprint(todos[i].ID) == todoId {
                if err := c.BodyParser(todo); err!= nil {
                    return c.Status(500).JSON(fiber.Map{"error": err.Error()})
                }
                todos[i].Completed = !todos[i].Completed
                return c.Status(200).JSON(fiber.Map{"msg": "updated successfully"})
            }
		}
		return c.Status(404).JSON(fiber.Map{"error": "unable to update"})
	})
	app.Delete("/api/tasks/:id", func(c *fiber.Ctx) error {
		taskId := c.Params("id")
		for i, todo  := range(todos){
			if fmt.Sprint(todo.ID) == taskId {
                todos = append(todos[:i], todos[i+1:]...)
                return c.Status(204).JSON(fiber.Map{"msg": "deleted successfully"})
            }
		}
		return c.Status(500).JSON("error", "could not delete")
	})

	log.Fatal(app.Listen(":8080"))
}