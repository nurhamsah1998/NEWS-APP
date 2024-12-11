package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nurhamsah1998/news/database"
	"github.com/nurhamsah1998/news/routes"
)

func main() {
	database.DBConnection()
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError
			fmt.Println(err, "<==")
			// Send custom error page
			return ctx.Status(code).SendString("Internal server errorrrrr")
		},
	})
	app.Use(recover.New())
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/user", routes.CreateUser)
	app.Get("/user", routes.GetUsers)
	app.Listen(":3000")
}
