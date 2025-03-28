package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello, World!!")

	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Body is required",
			})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)
		// *todo is a pointer to the todo struct
		// &todo is a pointer to the todo variable

		return c.Status(201).JSON(todo)

	})

	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = !todos[i].Completed
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"error": "Todo not found",
		})

	})

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				// append all the todos before and after the todo to be deleted
				// ... will spread the todos really similar to js
				// todos[:i] will get all the todos before the todo to be deleted
				// todos[i+1:] will get all the todos after the todo to be deleted
				// append will append all the todos before and after the todo to be deleted
				// todos = append(todos[:i], todos[i+1:]...)
		
				todos = append(todos[:i], todos[i+1:]...) 
				return c.Status(200).JSON(fiber.Map{ "success": true })
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	log.Fatal(app.Listen(":"+ PORT))

}
