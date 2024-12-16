package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/nurhamsah1998/news/database"
	"github.com/nurhamsah1998/news/routes"
)

func main() {
	/// ENV INITIALIZER
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	/// DB MIGRATION
	database.DBConnection()
	/// APP INITIALIZER
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError
			fmt.Println(err, "<==")
			if err != nil {
				return ctx.Status(code).JSON(fiber.Map{
					"message": "Iternal server error",
					"error":   err.Error(),
				})
			}
			// Send custom error page
			return ctx.Status(code).SendString("Internal server errorrrrr")
		},
	})
	app.Use(recover.New())
	routes.UserRoutes(app)
	routes.AuthRoutes(app)
	app.Listen(":3000")
}
