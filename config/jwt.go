package config

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	ID       uint   `json:"ID"`
	jwt.RegisteredClaims
}
