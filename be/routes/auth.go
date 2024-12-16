package routes

import (
	"strings"

	"github.com/gofiber/fiber/v2"
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
	errTransaction := database.Database.Db.Transaction(func(tx *gorm.DB) error {
		if errUser := tx.Create(&user).Error; errUser != nil {
			return errUser
		}
		if errProfile := tx.Create(&models.Profile{Fullname: form.FullName, UserID: int(user.ID)}).Error; errProfile != nil {
			return errProfile
		}
		return nil
	})
	/// EROR IF EMAIL ALREADY EXIST
	if strings.Contains(errTransaction.Error(), "duplicate key") {
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
	println(form.Email, "<EMAIL")
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
	return c.Status(200).JSON(fiber.Map{
		"message": "Successfully sign in",
	})
}

func AuthRoutes(app *fiber.App) {
	app.Post("/auth/sign-up", SignUp)
	app.Post("/auth/sign-in", SignIn)
}
