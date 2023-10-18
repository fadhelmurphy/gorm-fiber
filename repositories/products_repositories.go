package repositories

import (
	"go-pagination-gorm/models"

	"gorm.io/gorm"
)

type dbProduct struct {
	Conn *gorm.DB
}

// Create implements ProductRepository.
func (db *dbProduct) Create(product models.Product) error {
	return db.Conn.Create(&product).Error
}

// Delete implements ProductRepository.
func (db *dbProduct) Delete(idProduct int) error {
	return db.Conn.Delete(&models.Product{Id: idProduct}).Error
}

// GetAll implements ProductRepository.
func (db *dbProduct) GetAll() ([]models.Product, error) {
	var data []models.Product
	result := db.Conn.Find(&data)
	return data, result.Error
}

// GetById implements ProductRepository.
func (db *dbProduct) GetById(idProduct int) (models.Product, error) {
	var data models.Product
	result := db.Conn.Where("id", idProduct).First(&data)
	return data, result.Error
}

// Update implements ProductRepository.
func (db *dbProduct) Update(idProduct int, Product models.Product) error {
	return db.Conn.Where("id", idProduct).Updates(Product).Error
}

type ProductRepository interface {
	Create(Product models.Product) error
	Update(idProduct int, Product models.Product) error
	Delete(idProduct int) error
	GetById(idProduct int) (models.Product, error)
	GetAll() ([]models.Product, error)
}

func NewProductRepository(Conn *gorm.DB) ProductRepository {
	return &dbProduct{Conn: Conn}
}
