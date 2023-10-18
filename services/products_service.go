package services

import (
	"fmt"
	"go-pagination-gorm/helpers"
	"go-pagination-gorm/models"
	"go-pagination-gorm/repositories"

	"gorm.io/gorm"
)

type ProductService struct {
	productRepo repositories.ProductRepository
}

// Create implements ProductService.
func (service *ProductService) Create(products models.Product) helpers.Response {
	var response helpers.Response
	if err := service.productRepo.Create(products); err != nil {
		response.Status = 500
		response.Messages = "Failed to create a new product " + err.Error()
	} else {
		response.Status = 200
		response.Messages = "Success to create a new product"
	}
	return response
}

// Delete implements ProductService.
func (service *ProductService) Delete(idItem int) helpers.Response {
	var response helpers.Response
	if err := service.productRepo.Delete(idItem); err != nil {
		response.Status = 500
		response.Messages = fmt.Sprint("Failed to delete product : ", idItem)
	} else {
		response.Status = 200
		response.Messages = "Success to delete product"
	}
	return response
}

// GetAll implements ProductService.
func (service *ProductService) GetAll() helpers.Response {
	var response helpers.Response
	data, err := service.productRepo.GetAll()
	if err != nil {
		response.Status = 500
		response.Messages = "Failed to get products"
	} else {
		response.Status = 200
		response.Messages = "Success to get products"
		response.Data = data
	}
	return response
}

// GetById implements ProductService.
func (service *ProductService) GetById(idItem int) helpers.Response {
	var response helpers.Response
	data, err := service.productRepo.GetById(idItem)
	if err != nil {
		response.Status = 500
		response.Messages = fmt.Sprint("Failed to get product : ", idItem)
	} else {
		response.Status = 200
		response.Messages = "Success to get product"
		response.Data = data
	}
	return response
}

// Update implements ProductService.
func (service *ProductService) Update(idItem int, product models.Product) helpers.Response {
	var response helpers.Response
	if err := service.productRepo.Update(idItem, product); err != nil {
		response.Status = 500
		response.Messages = fmt.Sprint("Failed to update product ", idItem)
	} else {
		response.Status = 200
		response.Messages = "Success to update product"
	}
	return response
}

type IProductService interface {
	Create(product models.Product) helpers.Response
	Update(idItem int, product models.Product) helpers.Response
	Delete(idItem int) helpers.Response
	GetById(idItem int) helpers.Response
	GetAll() helpers.Response
}

func NewProductService(db *gorm.DB) ProductService {
	return ProductService{productRepo: repositories.NewProductRepository(db)}
}
