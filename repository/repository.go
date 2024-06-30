package repository

import (
	"autmtres/models"

	"gorm.io/gorm"
)

// ProductRepository interface defines the methods to be implemented
type ProductRepository interface {
	Create(product *models.Product) error
	GetAll() ([]models.Product, error)
	GetByID(id uint) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id uint) error
}

// GormProductRepository implements ProductRepository with GORM
type GormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) *GormProductRepository {
	return &GormProductRepository{db: db}
}

func (repo *GormProductRepository) Create(product *models.Product) error {
	return repo.db.Create(product).Error
}

func (repo *GormProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := repo.db.Find(&products).Error
	return products, err
}

func (repo *GormProductRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	err := repo.db.First(&product, id).Error
	return &product, err
}

func (repo *GormProductRepository) Update(product *models.Product) error {
	return repo.db.Save(product).Error
}

func (repo *GormProductRepository) Delete(id uint) error {
	return repo.db.Delete(&models.Product{}, id).Error
}

type OrderRepository interface {
	Create(order *models.Order) error
	GetAll() ([]models.Order, error)
	GetByID(id uint) (*models.Order, error)
	Update(order *models.Order) error
	Delete(id uint) error
}

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) *GormOrderRepository {
	return &GormOrderRepository{db: db}
}

func (repo *GormOrderRepository) Create(order *models.Order) error {
	return repo.db.Create(order).Error
}

func (repo *GormOrderRepository) GetAll() ([]models.Order, error) {
	var orders []models.Order
	err := repo.db.Preload("Products").Find(&orders).Error
	return orders, err
}

func (repo *GormOrderRepository) GetByID(id uint) (*models.Order, error) {
	var order models.Order
	err := repo.db.Preload("Products").First(&order, id).Error
	return &order, err
}

func (repo *GormOrderRepository) Update(order *models.Order) error {
	return repo.db.Save(order).Error
}

func (repo *GormOrderRepository) Delete(id uint) error {
	return repo.db.Delete(&models.Order{}, id).Error
}
