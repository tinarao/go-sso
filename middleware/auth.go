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

func Protected(fn fiber.Handler, roles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var headers map[string][]string
		headers = c.GetReqHeaders()
		if headers["Authorization"] != nil {
			checkedRole, err := JWTChecker(headers["Authorization"])
			if err != nil {
				log.Error(err)
				return c.Status(403).JSON(fiber.Map{
					"message": "Forbidden",
				})
			}

			var isValidRole bool
			for _, r := range roles {
				if checkedRole == r {
					isValidRole = true
				}
			}

			if !isValidRole {
				return c.Status(403).JSON(fiber.Map{
					"message": "Forbidden. For you.",
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

func JWTChecker(token []string) (string, error) {
	var arr = strings.Split(token[0], " ")
	if len(arr) != 2 {
		return "", errors.New("invalid token")
	}
	tokenStr := arr[1]

	parsedToken, err := jwt.ParseWithClaims(tokenStr, &config.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return "", err
	} else if token, ok := parsedToken.Claims.(*config.JwtClaims); ok {
		return token.Role, nil
	} else {
		return "", err
	}
}
