package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurhamsah1998/news/database"
	"github.com/nurhamsah1998/news/models"
	"github.com/nurhamsah1998/news/utils"
)

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	response := []models.UserResponse{}
	database.Database.Db.Find(&users)
	for _, value := range users {
		serializer := models.UserResponse{Id: int(value.ID), Email: value.Email}
		response = append(response, serializer)
	}

	return c.Status(200).JSON(utils.GlobalResponse{Data: response, Message: "success"})

}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	return c.Status(200).JSON(user)
}
