package handlers

import (
	"log"
	"tandigital/backend/database"
	"tandigital/backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Consumer models.Consumer

type CreateConsumerBody struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func RegisterConsumer(c *fiber.Ctx) error {
	var body CreateConsumerBody

	if err := c.BodyParser(&body); err != nil {
		return c.JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	var consumer Consumer
	consumer.Email = body.Email
	consumer.Password = body.Password

	database.DB.Create(&consumer)

	return c.JSON(fiber.Map{
		"message": "Account Created",
	})
}

func LoginConsumer(c *fiber.Ctx) error {

	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal error",
		})
	}

	var body CreateConsumerBody
	if err := c.BodyParser(&body); err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal error",
		})
	}

	var consumer Consumer
	database.DB.First(&consumer, "email = ?", body.Email)

	if consumer.Email != body.Email {
		return c.JSON(fiber.Map{
			"error": "Email doesn't exists",
		})
	}

	if consumer.Password != body.Password {
		return c.JSON(fiber.Map{
			"error": "Incorrect password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": consumer.Email,
		"role":  "consumer",
		"id":    consumer.ID,
	})

	tokenString, err := token.SignedString([]byte(envs["JWT_SECRET"]))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't login",
		})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}
