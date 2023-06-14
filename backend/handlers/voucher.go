package handlers

import (
	"log"
	"tandigital/backend/database"
	"tandigital/backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Voucher models.Voucher
type Transaction models.Transaction

type CreateVoucherBody struct {
	Name       string `json:"name" form:"name"`
	Code       string `json:"code" form:"code"`
	Value      int    `json:"value" form:"value"`
	Price      int    `json:"price" form:"price"`
	Quantity   int    `json:"quantity" form:"quantity"`
	ExpiryDate string `json:"expiry_date" form:"expiry_date"`
}

func GetAllVouchers(c *fiber.Ctx) error {
	var vouchers []Voucher

	d := database.DB.Find(&vouchers)
	if d.Error != nil {
		return c.JSON(fiber.Map{
			"error": d.Error,
		})
	}
	return c.JSON(vouchers)
}

func CreateVoucher(c *fiber.Ctx) error {
	role := c.Locals("role")

	if role != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	var body CreateVoucherBody

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02", body.ExpiryDate)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	var voucher Voucher
	voucher.Name = body.Name
	voucher.Code = body.Code
	voucher.Value = body.Value
	voucher.Price = body.Price
	voucher.Quantity = body.Quantity
	voucher.Status = "active"
	voucher.ExpiryDate = t

	log.Println(voucher)

	database.DB.Create(&voucher)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func UpdateVoucher(c *fiber.Ctx) error {
	role := c.Locals("role")

	if role != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	var body CreateVoucherBody

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	t, err := time.Parse("2006-01-02", body.ExpiryDate)
	if err != nil {
		return err
	}

	var voucher Voucher
	voucher.Name = body.Name
	voucher.Code = body.Code
	voucher.Value = body.Value
	voucher.Price = body.Price
	voucher.Quantity = body.Quantity
	voucher.ExpiryDate = t

	database.DB.Model(&voucher).Where("id = ?", c.Params("id")).Updates(&voucher)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func DeleteVoucher(c *fiber.Ctx) error {
	role := c.Locals("role")

	if role != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	var voucher Voucher
	if err := database.DB.First(&voucher, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Voucher not found",
		})
	}

	// database.DB.Model(&voucher).Where("id = ?", c.Params("id")).Update("status", "deleted")
	// database.DB.Delete(&voucher, c.Params("id"))

	if err := database.DB.Model(&voucher).Updates(map[string]interface{}{"status": "deleted"}).Error; err != nil {
		return err
	}

	if err := database.DB.Delete(&voucher).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func BuyVoucher(c *fiber.Ctx) error {
	role := c.Locals("role")

	if role != "consumer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	var voucher Voucher
	d := database.DB.First(&voucher, c.Params("id"))
	if d.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Voucher not found",
		})
	}

	if voucher.Quantity == 0 {
		return c.JSON(fiber.Map{
			"error": "Voucher is out of stock",
		})
	}

	voucher.Quantity = voucher.Quantity - 1

	database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&voucher).Updates(map[string]interface{}{"quantity": voucher.Quantity}).Error; err != nil {
			return err
		}

		var transaction Transaction
		transaction.ConsumerID = uint(c.Locals("id").(float64))
		transaction.Status = "paid"
		transaction.Vouchers = append(transaction.Vouchers, models.Voucher(voucher))

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		// d := database.DB.First(&voucher, c.Params("id"))
		// if d.Error != nil {
		// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		// 		"error": "Voucher not found",
		// 	})
		// }

		// // id := strconv.Itoa(int(voucher.ID))

		// if err := tx.Model(&transaction).Association("Vouchers").Append([]Voucher{}); err != nil {
		// 	return err
		// }

		return nil
	})
	// database.DB.Save(&voucher)

	// var transaction Transaction
	// transaction.ConsumerID = uint(c.Locals("id").(float64))
	// transaction.Status = "paid"

	// database.DB.Create(&transaction)

	// var transactionVoucher TransactionVoucher
	// transactionVoucher.TransactionID = transaction.ID
	// transactionVoucher.VoucherID = voucher.ID

	// database.DB.Create(&transactionVoucher)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func GetVoucherbyConsumer(c *fiber.Ctx) error {
	var vouchers []Voucher

	d := database.DB.Raw("select tv.transaction_id as id , v.name , v.price , v.code , v.expiry_date from transaction_voucher tv left join transactions t ON tv.transaction_id = t.id left join vouchers v ON tv.voucher_id = v.id where t.consumer_id = ?", c.Locals("id")).Scan(&vouchers)
	if d.Error != nil {
		return c.JSON(fiber.Map{
			"error": d.Error,
		})
	}

	return c.JSON(vouchers)
}
