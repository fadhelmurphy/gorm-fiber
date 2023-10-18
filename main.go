package main

import (
	"go-pagination-gorm/config"
	"github.com/gofiber/fiber/v2"
	"go-pagination-gorm/routes"
)

func main() {

	db := config.InitDB()

	app := fiber.New()
    server := routes.PublicRoutes(app, db)

    server.Listen(":3000")
}