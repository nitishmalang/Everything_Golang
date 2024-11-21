package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{}

func main() {
	app := fiber.New()

	app.Post("/users", func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}
		user.ID = uuid.New().String() 
		users = append(users, *user)
		return c.Status(fiber.StatusCreated).JSON(user)
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for _, user := range users {
			if user.ID == id {
				return c.JSON(user)
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	})

	app.Put("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, user := range users {
			if user.ID == id {
				if err := c.BodyParser(&users[i]); err != nil {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
				}
				users[i].ID = id // Ensure the ID is not overwritten
				return c.JSON(users[i])
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	})

	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, user := range users {
			if user.ID == id {
				users = append(users[:i], users[i+1:]...)
				return c.SendStatus(fiber.StatusNoContent)
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	})

	app.Listen(":3000")
}
