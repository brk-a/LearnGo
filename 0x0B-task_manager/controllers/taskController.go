package controllers

import (
	"context"
	"time"

	"github.com/brk-a/task_manager/database"
	"github.com/brk-a/task_manager/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var taskCollection *mongo.Collection = database.GetCollection(database.Client, "task_collection")
var validate = validator.New()

// GetTasks retrieves all tasks from the database and returns them as a JSON response.
//
// The function uses a context with a timeout of 100 seconds to ensure that the operation does not take too long.
// It initialises an empty slice of Task structs to store the retrieved tasks.
//
// The function then uses the Find method from the mongo-driver to query the database for all tasks.
// If an error occurs during the query, the function returns a 500 Internal Server Error status code with an error message.
//
// After retrieving the tasks, the function uses the All method to decode the results into the allTasks slice.
// If an error occurs during the decoding, the function returns a 500 Internal Server Error status code with an error message.
//
// Finally, the function returns a 200 OK status code with the first task from the allTasks slice as the JSON response.
func GetTasks(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	var allTasks []models.Task

	result, err := taskCollection.Find(ctx, bson.M{})
	defer cancel()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not retrieve tasks",
		})
	}
	if err = result.All(ctx, &allTasks); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error decoding tasks",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": allTasks[0],
	})
}

// GetTask retrieves a single task from the database based on the provided ID.
//
// The function uses a context with a timeout of 100 seconds to ensure that the operation does not take too long.
// It then queries the database for a task with the matching ID.
// If an error occurs during the query, the function returns a 404 Not Found status code with an error message.
//
// After retrieving the task, the function returns a 200 OK status code with the task as the JSON response.
func GetTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	id := c.Params("id")
	var task models.Task

	err := taskCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	defer cancel()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "task not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": task,
	})
}

// CreateTask creates a new task in the database based on the provided task data.
//
// The function uses a context with a timeout of 100 seconds to ensure that the operation does not take too long.
// It then parses the request body to extract the task data.
// If an error occurs during the parsing, the function returns a 400 Bad Request status code with an error message.
//
// After validating the task data, the function sets the default values for the task.
// It then inserts the task into the database.
// If an error occurs during the insertion, the function returns a 500 Internal Server Error status code with an error message.
//
// Finally, the function returns a 201 Created status code with the inserted task as the JSON response.
func CreateTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	var task models.Task

	defer cancel()
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}
	if err := validate.Struct(task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot validate task",
		})
	}
	if task.Body == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "task body cannot be empty",
		})
	}
	if task.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "task title cannot be empty",
		})
	}

	task.Completed = false
	task.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	task.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	task.ID = primitive.NewObjectID()
	task.Task_id = task.ID.Hex()
	
	_, err := taskCollection.InsertOne(ctx, task)
	defer cancel()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error creating task",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": "task created successfully",
	})
}

// UpdateTask updates an existing task in the database based on the provided ID and task data.
//
// The function uses a context with a timeout of 100 seconds to ensure that the operation does not take too long.
// It then parses the request body to extract the task data.
// If an error occurs during the parsing, the function returns a 400 Bad Request status code with an error message.
//
// After validating the task data, the function constructs an update object with the provided fields.
// It then updates the task in the database.
// If an error occurs during the update, the function returns a 500 Internal Server Error status code with an error message.
//
// Finally, the function returns a 200 OK status code with a success message.
func UpdateTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	id := c.Params("id")
	var task models.Task

	defer cancel()
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	var updateObj primitive.D
	filter := bson.M{"_id": id}
	if task.Title != "" {
		updateObj = append(updateObj, bson.E{Key: "title", Value: task.Title})
	}
	if task.Body != "" {
		updateObj = append(updateObj, bson.E{Key: "body", Value: task.Body})
	}
	updateObj = append(updateObj, bson.E{Key: "completed", Value: task.Completed})
	task.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: task.Updated_at})

	upsert := true
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	_, err := taskCollection.UpdateOne(ctx, filter, bson.M{"$set": updateObj}, &opt)
	defer cancel()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error updating task",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "task updated successfully",
	})
}

// DeleteTask deletes an existing task from the database based on the provided ID.
//
// The function uses a context with a timeout of 100 seconds to ensure that the operation does not take too long.
// It then queries the database for a task with the matching ID and deletes it.
// If an error occurs during the deletion, the function returns a 500 Internal Server Error status code with an error message.
//
// Finally, the function returns a 200 OK status code with a success message.
func DeleteTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	id := c.Params("id")

	filter := bson.M{"_id": id}
	_, err := taskCollection.DeleteOne(ctx, filter)
	defer cancel()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error deleting task",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "task deleted successfully",
	})
}
