package routes

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/nurhamsah1998/news/database"
	"github.com/nurhamsah1998/news/middleware"
	"github.com/nurhamsah1998/news/models"
	"github.com/nurhamsah1998/news/utils"
)

func GetAllPost(c *fiber.Ctx) error {
	posts := []models.NewsPost{}
	var totalData int64
	queryPage, queryLimit := c.QueryInt("page"), (c.QueryInt("limit"))
	if queryLimit == 0 {
		queryLimit = 10
	}
	if queryPage == 0 {
		queryPage = 1
	}
	offside := (queryPage - 1) * queryLimit
	database.Database.Db.Model(&models.NewsCategory{}).Count(&totalData)
	totalPage := math.Ceil(float64(totalData / int64(queryLimit)))
	database.Database.Db.Limit(queryLimit).Offset(offside).Find(&posts)
	return c.Status(200).JSON(utils.GlobalResponse{
		Data:    posts,
		Message: "success",
		Meta: utils.MetaData{TotalPage: int(totalPage),
			Page: queryPage, TotalData: totalData}})
}

func PostRoutes(app *fiber.App) {
	postApi := app.Group("/news-post", middleware.UserMiddleware)
	postApi.Get("", GetAllPost)
}
