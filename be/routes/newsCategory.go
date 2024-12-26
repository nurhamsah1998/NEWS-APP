package routes

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/nurhamsah1998/news/database"
	"github.com/nurhamsah1998/news/middleware"
	"github.com/nurhamsah1998/news/models"
	"github.com/nurhamsah1998/news/utils"
)

func GetAllCategory(c *fiber.Ctx) error {
	categories := []models.NewsCategory{}
	var totalData int64
	queryPage, queryLimit := c.QueryInt("page"), c.QueryInt("limit")
	if queryLimit == 0 {
		queryLimit = 10
	}
	if queryPage == 0 {
		queryPage = 1
	}
	offside := (queryPage - 1) * queryLimit
	database.Database.Db.Model(&models.NewsCategory{}).Count(&totalData)
	totalPage := math.Ceil(float64(totalData / int64(queryLimit)))
	database.Database.Db.Limit(queryLimit).Offset(offside).Find(&categories)
	return c.Status(200).JSON(utils.GlobalResponse{
		Data:    categories,
		Message: "success",
		Meta: utils.MetaData{TotalPage: int(totalPage),
			Page: queryPage, TotalData: totalData}})
}

func PostCategory(c *fiber.Ctx) error {
	var category models.NewsCategory
	err := c.BodyParser(&category)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Something wrong",
		})
	}
	if category.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Category name cannot be empty",
		})
	}
	result := database.Database.Db.Save(&category)
	if result.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Something wrong, Cannot create category",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Successfully create category",
	})
}
func DeleteCategory(c *fiber.Ctx) error {
	idCategory := c.Params("id")
	var category models.NewsCategory
	result := database.Database.Db.Delete(&category, idCategory)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("Delete category failed")
	}
	return c.Status(200).SendString("Success")
}
func CategoryRoutes(app *fiber.App) {
	categoryApi := app.Group("/news-category", middleware.UserMiddleware)
	categoryApi.Get("", GetAllCategory)
	categoryApi.Post("", PostCategory)
	categoryApi.Delete("/:id", DeleteCategory)
}
