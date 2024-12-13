package routes

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/nurhamsah1998/news/database"
	"github.com/nurhamsah1998/news/models"
	"github.com/nurhamsah1998/news/utils"
)

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	var totalData int64
	response := []models.UserResponse{}
	queryPage, queryLimit := c.QueryInt("page"), c.QueryInt("limit")
	if queryLimit == 0 {
		queryLimit = 10
	}
	if queryPage == 0 {
		queryPage = 1
	}
	offside := (queryPage - 1) * queryLimit
	database.Database.Db.Model(&models.User{}).Count(&totalData)
	totalPage := math.Ceil(float64(totalData / int64(queryLimit)))

	database.Database.Db.Limit(queryLimit).Offset(offside).Find(&users)
	for _, value := range users {
		serializer := models.UserResponse{Id: int(value.ID), Email: value.Email}
		response = append(response, serializer)
	}

	return c.Status(200).JSON(utils.GlobalResponse{Data: response, Message: "success", Meta: utils.MetaData{TotalPage: int(totalPage), Page: queryPage, TotalData: totalData}})

}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	return c.Status(200).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	idUser := c.Params("id")
	var user models.User
	database.Database.Db.Delete(&user, idUser)
	println(user.Email, "<===")
	return c.Status(200).SendString("Success")
}

func EditUserById(c *fiber.Ctx) error {
	idUser, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("The id must be a number!")
	}
	var user models.User
	result := database.Database.Db.Find(&user, idUser)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("user not found!")
	}
	type DataBody struct {
		Email string
	}
	var updateBody DataBody
	if err := c.BodyParser(&updateBody); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	resultAfterEdit := database.Database.Db.Save(&updateBody)
	if resultAfterEdit.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("Failed change user")
	}
	return c.Status(200).JSON(updateBody)
}

func GetUserById(c *fiber.Ctx) error {
	idUser, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("The id must be a number!")
	}
	var user models.User
	result := database.Database.Db.Find(&user, idUser)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("user not found!")
	}
	return c.Status(200).JSON(&utils.GlobalResponse{Data: user, Message: "Success"})
}

func UserRoutes(app *fiber.App) {
	app.Post("/user", CreateUser)
	app.Get("/user", GetUsers)
	app.Delete("/user/:id", DeleteUser)
	app.Get("/user/:id", GetUserById)
	app.Patch("/user/:id", EditUserById)
}
