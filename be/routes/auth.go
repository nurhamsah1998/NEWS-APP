package routes

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nurhamsah1998/news/database"
	"github.com/nurhamsah1998/news/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SignUpForm struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(c *fiber.Ctx) error {
	var form SignUpForm
	var user models.User
	if err := c.BodyParser(&form); err != nil {
		c.Status(400).JSON(err.Error())
	}
	if len(form.Password) < 8 {
		return c.Status(500).JSON(fiber.Map{
			"message": "Password must be not under 8 character",
		})
	}
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(form.Password), 10)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Intenal server error",
		})
	}
	user.Password = string(hashPwd)
	user.Email = form.Email
	var errorString string
	database.Database.Db.Transaction(func(tx *gorm.DB) error {
		if errUser := tx.Create(&user).Error; errUser != nil {
			errorString = errUser.Error()
			return errUser
		}
		if errProfile := tx.Create(&models.Profile{Fullname: form.FullName, UserID: int(user.ID)}).Error; errProfile != nil {
			errorString = errProfile.Error()
			return errProfile
		}
		return nil
	})
	/// EROR IF EMAIL ALREADY EXIST
	if strings.Contains(errorString, "duplicate key") {
		errorString = ""
		return c.Status(400).JSON(fiber.Map{
			"message": form.Email + " Already exist. try something else",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Successfully sign up",
	})
}

func SignIn(c *fiber.Ctx) error {
	var user models.User
	var form SignUpForm
	if err := c.BodyParser(&form); err != nil {
		c.Status(400).JSON(err.Error())
	}
	database.Database.Db.First(&user, "email = ?", form.Email)
	if user.ID == 0 {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid crendetial",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid email or password.",
		})
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":     user.ID,
		"expired": time.Now().Add(time.Hour * 8).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Cannot sign in token",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Successfully sign in",
		"data": struct {
			AccessToken string `json:"access_token"`
		}{
			AccessToken: tokenString,
		},
	})
}

func AuthRoutes(app *fiber.App) {
	app.Post("/auth/sign-up", SignUp)
	app.Post("/auth/sign-in", SignIn)
}
