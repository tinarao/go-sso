package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v3/log"
	"github.com/tinarao/go-sso/models"
)

func Login(c *fiber.Ctx) error {
	doc := &models.LoginDTO{}
	err := c.BodyParser(&doc)
	if err != nil {
		log.Error(err)
		return err
	}

	return c.Status(201).JSON(fiber.Map{
		"credentials": doc,
	})
}

func Register(c *fiber.Ctx) error {
	return c.SendString("Register")
}
