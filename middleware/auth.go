package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tinarao/go-sso/config"
	"os"
	"strings"
)

type Cookie struct {
	AccessToken string `cookie:"access_token"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Protected(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var headers map[string][]string
		headers = c.GetReqHeaders()
		if headers["Authorization"] != nil {
			if err := JWTChecker(headers["Authorization"]); err != nil {
				log.Error(err)
				return c.Status(403).JSON(fiber.Map{
					"message": "Forbidden",
				})
			}
		} else {
			return c.Status(401).JSON(fiber.Map{
				"message": "Forbidden",
			})
		}
		return fn(c)
	}
}

func JWTChecker(token []string) error {
	var arr = strings.Split(token[0], " ")
	if len(arr) != 2 {
		return errors.New("Invalid token")
	}
	tokenStr := arr[1]

	parsedToken, err := jwt.ParseWithClaims(tokenStr, &config.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return err
	} else if _, ok := parsedToken.Claims.(*config.JwtClaims); ok {
		return nil
	} else {
		return err
	}
}
