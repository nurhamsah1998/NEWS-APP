package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurhamsah1998/news/database"
	"github.com/nurhamsah1998/news/models"
)

func CreateUser(c *fiber.Ctx) error {
	var user *models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	return c.Status(200).JSON(user)
}
