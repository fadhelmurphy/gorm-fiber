package controllers

import (
	"go-pagination-gorm/models"
	"go-pagination-gorm/services"
	"net/http"
	"strconv"

	vl "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductController struct {
	ProductService services.ProductService
	validate       vl.Validate
}

func (controller ProductController) Create(c *fiber.Ctx) error {

	type payload struct {
		Id          uint   `json:"id"`
		Title       string `json:"title" validate:"required"`
		Description string `json:"description" validate:"required"`
		Image       string `json:"image"`
		Price       int    `json:"price" validate:"required"`
	}
	payloadValidator := new(payload)
	if err := controller.validate.Struct(payloadValidator); err != nil {
		return err
	}

	result := controller.ProductService.Create(models.Product{
		Title:       payloadValidator.Title,
		Description: payloadValidator.Description,
		Image:       payloadValidator.Image,
		Price:       payloadValidator.Price,
	})

	return c.Status(http.StatusOK).JSON(result)
}
func (controller ProductController) Update(c *fiber.Ctx) error {
	type payload struct {
		Id          uint   `json:"id"`
		Title       string `json:"title" validate:"required"`
		Description string `json:"description" validate:"required"`
		Image       string `json:"image"`
		Price       int    `json:"price" validate:"required"`
	}

	payloadValidator := new(payload)

	idItem, _ := strconv.Atoi(c.Params("id"))
	result := controller.ProductService.Update(idItem,
		models.Product{
			Title:       payloadValidator.Title,
			Description: payloadValidator.Description,
			Image:       payloadValidator.Image,
			Price:       payloadValidator.Price,
		},
	)

	return c.Status(http.StatusOK).JSON(result)
}

func (controller ProductController) Delete(c *fiber.Ctx) error {
	idItem, _ := strconv.Atoi(c.Params("id"))
	result := controller.ProductService.Delete(idItem)

	return c.Status(http.StatusOK).JSON(result)
}

func (controller ProductController) GetAll(c *fiber.Ctx) error {
	result := controller.ProductService.GetAll()
	return c.Status(http.StatusOK).JSON(result)
}

func (controller ProductController) GetById(c *fiber.Ctx) error {
	idItem, _ := strconv.Atoi(c.Params("id"))
	result := controller.ProductService.GetById(idItem)
	return c.Status(http.StatusOK).JSON(result)
}

func NewProductController(db *gorm.DB) ProductController {
	service := services.NewProductService(db)
	controller := ProductController{
		ProductService: service,
	}

	return controller
}
