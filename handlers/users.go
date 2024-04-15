package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tinarao/go-sso/db"
	"github.com/tinarao/go-sso/models"
)

func GetUsers(c *fiber.Ctx) error {
	data := []models.User{}
	db.DB.Db.Find(&data)

	if len(data) == 0 {
		return c.Status(200).JSON(fiber.Map{
			"message": "Table is empty!",
			"users":   data,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"amount of users": len(data),
		"users":           data,
	})
}
