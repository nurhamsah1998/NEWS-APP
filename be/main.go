package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurhamsah1998/news/database"
	"github.com/nurhamsah1998/news/routes"
)

func main() {
	database.DBConnection()
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/user", routes.CreateUser)
	app.Get("/user", routes.GetUsers)

	app.Listen(":3000")
}
