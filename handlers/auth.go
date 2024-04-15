package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v3/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tinarao/go-sso/config"
	"github.com/tinarao/go-sso/db"
	"github.com/tinarao/go-sso/models"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func Login(c *fiber.Ctx) error {
	doc := &models.LoginDTO{}
	err := c.BodyParser(&doc)
	if err != nil {
		log.Error(err)
		return err
	}

	candidate := &models.User{}
	db.DB.Db.Where("email = ?", doc.Email).First(&candidate)

	if candidate.Email != doc.Email {
		return c.Status(400).JSON(fiber.Map{
			"message": "User with this credentials does not exist",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(candidate.Password), []byte(doc.Password)); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Wrong credentials",
		})
	}

	claims := config.JwtClaims{
		candidate.Username,
		candidate.Role,
		candidate.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Error(err)
		return err
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Successfully logged in!",
		"token":   signedString,
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

	existingUser := &models.User{}
	db.DB.Db.Where("email = ?", doc.Email).First(&existingUser)

	if existingUser.Email == doc.Email {
		return c.Status(400).JSON(fiber.Map{
			"message": "Users with this credentials is already created",
		})
	}

	if doc.Role != "" {
		return c.Status(403).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}

	doc.Role = "user"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(doc.Password), 10)
	if err != nil {
		log.Error(err)
		return err
	}

	doc.Password = string(hashedPassword)
	db.DB.Db.Create(&doc)

	return c.Status(201).JSON(fiber.Map{
		"message": "Successfully created an account!",
	})
}
