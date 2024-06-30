package services

import (
	"autmtres/models"
	services "autmtres/repository"
)

// ProductService interface defines the methods to be implemented
type ProductService interface {
	CreateProduct(product *models.Product) error
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
}

// ProductServiceImpl implements ProductService
type ProductServiceImpl struct {
	repo services.ProductRepository
}

func NewProductService(repo services.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{repo: repo}
}

func (service *ProductServiceImpl) CreateProduct(product *models.Product) error {
	return service.repo.Create(product)
}

func (service *ProductServiceImpl) GetAllProducts() ([]models.Product, error) {
	return service.repo.GetAll()
}

func (service *ProductServiceImpl) GetProductByID(id uint) (*models.Product, error) {
	return service.repo.GetByID(id)
}

func (service *ProductServiceImpl) UpdateProduct(product *models.Product) error {
	return service.repo.Update(product)
}

func (service *ProductServiceImpl) DeleteProduct(id uint) error {
	return service.repo.Delete(id)
}
