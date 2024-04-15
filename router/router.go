package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tinarao/go-sso/handlers"
	"github.com/tinarao/go-sso/middleware"
)

var moderator = []string{"moderator", "admin"}
var admin = []string{"admin"}

func SetupRoutes(app *fiber.App) {

	auth := app.Group("/auth")
	users := app.Group("/users")

	auth.Post("/login", handlers.Login)
	auth.Post("/register", handlers.Register)

	users.Get("/", middleware.Protected(handlers.GetUsers, moderator))
	users.Get("/role/:email", handlers.GetUserRole)
	users.Get("/info/:email", handlers.GetUserInfo)
	users.Delete("/:email/:token", handlers.DeleteUser)
	users.Delete("/all", middleware.Protected(handlers.DeleteAllUsers, admin))
}
