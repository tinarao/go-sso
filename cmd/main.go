package main

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/tinarao/go-sso/db"
	"github.com/tinarao/go-sso/router"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env: ", err)
	}

	app := fiber.New()
	port := os.Getenv("PORT")

	app.Use(logger.New())
	app.Use(swagger.New(swagger.Config{
		BasePath: "/docs",
		Path:     "docs",
		Title:    "Go SSO docs",
	}))

	router.SetupRoutes(app)
	db.Connect()

	app.Listen(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
