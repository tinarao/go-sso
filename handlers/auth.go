package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v3/log"
	"github.com/tinarao/go-sso/db"
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
	doc := &models.User{}
	err := c.BodyParser(&doc)
	if err != nil {
		log.Error(err)
		return err
	}

	if (len(doc.Username) == 0) || (len(doc.Email) == 0) || (len(doc.Password) == 0) {
		return c.Status(400).JSON(fiber.Map{
			"message": "Wrong request: not all nesessary fields are provided",
		})
	}

	var existingUser models.User
	db.DB.Db.Where("email = ?", doc.Email).First(&existingUser)

	if existingUser.Email == doc.Email {
		return c.Status(400).JSON(fiber.Map{
			"message": "Users with this credentials is already created",
		})
	}

	var validRole bool

	if doc.Role == "" {
		doc.Role = "user"
	}

	for _, x := range models.Roles {
		if x == doc.Role {
			validRole = true
		}
	}

	if !validRole {
		return c.Status(400).JSON(fiber.Map{
			"message":   "Provided role is invalid",
			"validRole": validRole,
		})
	}

	db.DB.Db.Create(&doc)

	return c.Status(201).JSON(fiber.Map{
		"message": "Successfully created an account",
		"user":    doc,
	})
}
