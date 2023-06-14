package middleware

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func CheckAdminToken(c *fiber.Ctx) error {
	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secret_key := envs["JWT_SECRET"]

	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "No token provided",
		})
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		return []byte(secret_key), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// return c.JSON(claims)
		role := claims["role"].(string)
		c.Locals("role", role)
		// return c.Send([]byte(claims["role"].(string)))
		return c.Next()
	} else {
		return c.JSON(fiber.Map{
			"error": err,
		})
	}
}

func CheckConsumerToken(c *fiber.Ctx) error {
	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secret_key := envs["JWT_SECRET"]

	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "No token provided",
		})
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		return []byte(secret_key), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// return c.JSON(claims)
		role := claims["role"].(string)
		id := claims["id"].(float64)

		c.Locals("role", role)
		c.Locals("id", id)

		// return c.Send([]byte(claims["role"].(string)))
		return c.Next()
	} else {
		return c.JSON(fiber.Map{
			"error": err,
		})
	}
}
