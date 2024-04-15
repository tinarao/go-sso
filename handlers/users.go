package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v3/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tinarao/go-sso/config"
	"github.com/tinarao/go-sso/db"
	"github.com/tinarao/go-sso/models"
	"os"
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

func GetUserRole(c *fiber.Ctx) error {
	paramEmail := c.Params("email")
	user := &models.User{}

	db.DB.Db.Where("email = ?", paramEmail).First(&user)
	isAdmin := user.Role == "admin"

	if user.Email != paramEmail {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email is not provided or user is not registered",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"email":   paramEmail,
		"isAdmin": isAdmin,
	})
}

func GetUserInfo(c *fiber.Ctx) error {
	paramEmail := c.Params("email")
	user := &models.User{}

	db.DB.Db.Where("email = ?", paramEmail).First(&user)
	if user.Email != paramEmail {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email is not provided or user is not registered",
		})
	}

	userInfo := &models.UserInfoDTO{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		Username:  user.Username,
		Role:      user.Role,
	}

	return c.Status(200).JSON(fiber.Map{
		"user": userInfo,
	})

}

func DeleteUser(c *fiber.Ctx) error {
	emailParam := c.Params("email")
	tokenParam := c.Params("token")

	tokenClaims := &config.JwtClaims{}

	token, err := jwt.ParseWithClaims(tokenParam, &config.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Debug(err)
		log.Debugf("\ntoken: %v\n", tokenParam)
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid token provided",
		})
	} else if claims, ok := token.Claims.(*config.JwtClaims); ok {
		tokenClaims = claims
	} else {
		log.Error("unknown claims type, cannot proceed")
	}

	userToDelete := &models.User{}
	requester := &models.User{}
	db.DB.Db.Where("email = ?", emailParam).First(&userToDelete)
	db.DB.Db.Where("ID = ?", tokenClaims.ID).First(&requester)

	if requester.Role != "admin" && userToDelete.ID != requester.ID {
		return c.Status(403).JSON(fiber.Map{
			"message": "You can only delete your own account",
		})
	} else if userToDelete.Role == "admin" {
		return c.Status(403).JSON(fiber.Map{
			"message": "You can't delete this account",
		})
	}

	db.DB.Db.Where("email = ?", emailParam).Delete(&userToDelete)

	return c.Status(200).JSON(fiber.Map{
		"deleted user": userToDelete.Email,
	})
}
