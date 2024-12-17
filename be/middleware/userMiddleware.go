package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nurhamsah1998/news/database"
	"github.com/nurhamsah1998/news/models"
)

func UserMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Headers Authorization required!",
		})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		hmacSampleSecret := []byte(os.Getenv("SECRET_TOKEN"))
		return hmacSampleSecret, nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["expired"].(float64) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Token expired",
			})
		}
		var user models.User
		database.Database.Db.Find(&user, claims["sub"])
		if user.ID == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorization",
			})
		}
		return c.Next()
	} else {
		return err
	}
}
