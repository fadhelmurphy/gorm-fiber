package provider

import (
	"go-pagination-gorm/config"
	"go-pagination-gorm/domain/products/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func ProvideService()( app *fiber.App) {
    db := config.InitDB()
    
    app = fiber.New()
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