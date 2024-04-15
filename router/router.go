package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tinarao/go-sso/handlers"
	"github.com/tinarao/go-sso/middleware"
)

func SetupRoutes(app *fiber.App) {

	auth := app.Group("/auth")
	users := app.Group("/users")

	auth.Post("/login", handlers.Login)
	auth.Post("/register", handlers.Register)

	users.Get("/", middleware.Protected(handlers.GetUsers))
	users.Get("/role/:email", handlers.GetUserRole)
	users.Get("/info/:email", handlers.GetUserInfo)
	users.Delete("/:email/:token", middleware.Protected(handlers.DeleteUser))
}
