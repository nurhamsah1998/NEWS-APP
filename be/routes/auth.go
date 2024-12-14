package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type SignUpForm struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(c *fiber.Ctx) error {
	var form SignUpForm
	if err := c.BodyParser(&form); err != nil {
		c.Status(400).JSON(err.Error())
	}

	fmt.Println(form.Email, "<===")
	return c.Status(200).JSON(fiber.Map{
		"message": "Successfully sign up",
		"data":    form,
	})
}

func AuthRoutes(app *fiber.App) {
	app.Post("/auth/signup", SignUp)
}
