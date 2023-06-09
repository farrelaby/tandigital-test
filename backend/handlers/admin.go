package handlers

import (
	"log"
	"tandigital/backend/database"
	"tandigital/backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Admin models.Admin

type CreateAdminBody struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func GetAllAdmins(c *fiber.Ctx) error {
	var admins []Admin

	database.DB.Find(&admins)

	return c.JSON(admins)
}

func CreateAdmin(c *fiber.Ctx) error {

	var body CreateAdminBody
	// if err := c.BodyParser(&body); err != nil {
	// 	c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 		"error": "Cannot parse JSON",
	// 	})
	// 	return err
	// }

	if err := c.BodyParser(&body); err != nil {
		log.Fatal(err)
	}

	// hashedpass, err := utils.HashPassword(body.Password)
	// if err != nil {
	// 	c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 		"error": "invalid goblok",
	// 	})

	// 	return err
	// }

	var admin Admin
	admin.Email = body.Email
	admin.Password = body.Password

	database.DB.Create(&admin)

	return c.JSON(&admin)
}

func LoginAdmin(c *fiber.Ctx) error {
	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var body CreateAdminBody
	if err := c.BodyParser(&body); err != nil {
		log.Fatal(err)
	}

	var admin Admin
	database.DB.First(&admin, "email = ?", body.Email)

	if admin.Email != body.Email {
		return c.JSON(fiber.Map{
			"error": "Email doesn't exists",
		})
	}

	if admin.Password != body.Password {
		return c.JSON(fiber.Map{
			"error": "Incorrect password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": admin.Email,
		"role":  "admin",
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
