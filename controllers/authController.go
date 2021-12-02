package controllers

import (
	"time"
	"api-gofiber/test/models"
	"api-gofiber/test/helpers"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	database := c.Locals("DB").(*gorm.DB);	

	resultUser := map[string]interface{}{}

	queryUser := database.Model(&models.User{})
		queryUser.Select("id","password","email")
		queryUser.Where("email = ?", user)
		queryUser.First(&resultUser)

	if len(resultUser) == 0 {
		return c.Status(200).JSON(fiber.Map{
			"message" : "Email tidak ditemukan",
		})
	}
	
	var isValidPassword bool = helpers.CheckPasswordHash(
		pass,
		resultUser["password"].(string),
	)

	if isValidPassword == false {
		return c.Status(200).JSON(fiber.Map{
			"message" : "Password Salah",
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func Me(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}

func Logout(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message" : "Logout",
	})
}

func Signup(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message" : "Signup",
	})
}

func RefreshToken(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message" : "Refresh Token",
	})
}