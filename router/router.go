package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tinarao/go-sso/handlers"
)

func SetupRoutes(app *fiber.App) {

	auth := app.Group("/auth")

	auth.Post("/login", handlers.Login)
	auth.Post("/register", handlers.Register)

}
