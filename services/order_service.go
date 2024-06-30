package services

import (
	"autmtres/models"
	"autmtres/repository"
)

// OrderService interface defines the methods to be implemented
type OrderService interface {
	CreateOrder(order *models.Order) error
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(id uint) (*models.Order, error)
	UpdateOrder(order *models.Order) error
	DeleteOrder(id uint) error
}

// OrderServiceImpl implements OrderService
type OrderServiceImpl struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{repo: repo}
}

func (service *OrderServiceImpl) CreateOrder(order *models.Order) error {
	return service.repo.Create(order)
}

func (service *OrderServiceImpl) GetAllOrders() ([]models.Order, error) {
	return service.repo.GetAll()
}

func (service *OrderServiceImpl) GetOrderByID(id uint) (*models.Order, error) {
	return service.repo.GetByID(id)
}

func (service *OrderServiceImpl) UpdateOrder(order *models.Order) error {
	return service.repo.Update(order)
}

func (service *OrderServiceImpl) DeleteOrder(id uint) error {
	return service.repo.Delete(id)
}
