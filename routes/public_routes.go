package routes

import (
	"go-pagination-gorm/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func PublicRoutes(app *fiber.App, db *gorm.DB) *fiber.App {
	apiV1 := app.Group("api/v1")
	apiV1.Use(cors.New())

	productController := controllers.NewProductController(db)
	apiV1.Post("products", productController.Create)
	apiV1.Put("products/:id", productController.Update)
	apiV1.Delete("products/:id", productController.Delete)
	apiV1.Get("products", productController.GetAll)
	apiV1.Get("products/:id", productController.GetById)

	return app
}
