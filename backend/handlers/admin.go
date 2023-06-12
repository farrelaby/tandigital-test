package handlers

import (
	"log"
	"tandigital/backend/db"
	"tandigital/backend/models"

	"github.com/gofiber/fiber/v2"
)

type Admin models.Admin

type CreateAdminBody struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func GetAllAdmins(c *fiber.Ctx) error {
	var admins []Admin

	db.Init().Find(&admins)
	return c.JSON(admins)
}

func CreateAdmin(c *fiber.Ctx) error {

	var body CreateAdminBody
	if err := c.BodyParser(body); err != nil {
		return c.JSON(err)
	}

	var admin Admin
	admin.Email = body.Email
	admin.Password = body.Password
	db.Init().Create(&admin)
	return c.JSON(admin)

}

func TestBody(c *fiber.Ctx) error {

	// var body CreateAdminBody
	// c.BodyParser(&body)

	body := new(CreateAdminBody)

	// create an error handler when the request body param is not the same as the struct
	// if err := c.BodyParser(body); err != nil {
	// 	return c.JSON(err)
	// }

	// create an error handler when parsing the body
	if err := c.BodyParser(body); err != nil {
		return c.JSON(err)
	}

	log.Println(body.Email)
	return c.JSON(body)
}
