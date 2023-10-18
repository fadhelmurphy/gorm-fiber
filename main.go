package main

import (
	"go-pagination-gorm/config"
	"github.com/gofiber/fiber/v2"
	"go-pagination-gorm/routes"
)

func main() {

	db := config.InitDB()

	app := fiber.New()
    routes.PublicRoutes(app, db)

    app.Listen(":3000")
}