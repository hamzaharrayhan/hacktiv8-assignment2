package service

import (
	"assignment2/models"
	"assignment2/payload"
	"assignment2/repository"
	"strconv"
)

type Service interface {
	GetOrders() ([]models.Orders, error)
	CreateOrder(orderInput payload.OrderPayload) (models.Orders, error)
	UpdateOrder(orderID string, orderInput payload.OrderPayload) (models.Orders, error)
	DeleteOrder(orderID string) error
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) GetOrders() ([]models.Orders, error) {
	orders, err := s.repository.GetOrders()
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (s *service) CreateOrder(orderInput payload.OrderPayload) (models.Orders, error) {
	order := models.Orders{
		CustomerName: orderInput.CustomerName,
		OrderedAt:    orderInput.OrderedAt,
	}

	if len(orderInput.Items) > 0 {
		order.Items = orderInput.Items
	}
	orders, err := s.repository.CreateOrder(order)
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (s *service) UpdateOrder(orderID string, orderInput payload.OrderPayload) (models.Orders, error) {
	id, err := strconv.Atoi(orderID)
	if err != nil {
		return models.Orders{}, err
	}
	order := models.Orders{
		ID:           uint(id),
		CustomerName: orderInput.CustomerName,
		OrderedAt:    orderInput.OrderedAt,
	}

	if len(orderInput.Items) > 0 {
		order.Items = orderInput.Items
	}

	orders, err := s.repository.UpdateOrder(order)
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (s *service) DeleteOrder(orderID string) error {
	id, err := strconv.Atoi(orderID)
	if err != nil {
		return err
	}

	err = s.repository.DeleteOrder(uint(id))
	if err != nil {
		return err
	}

	return nil
}
