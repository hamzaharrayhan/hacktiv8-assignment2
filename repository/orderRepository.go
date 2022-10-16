package repository

import (
	"assignment2/models"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	CreateOrder(order models.Orders) (models.Orders, error)
	GetOrders() ([]models.Orders, error)
	UpdateOrder(order models.Orders) (models.Orders, error)
	DeleteOrder(orderID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetOrders() ([]models.Orders, error) {
	var orders []models.Orders
	err := r.db.Preload("Items").Find(&orders).Error
	if err != nil {
		fmt.Println("Error getting orders data:", err)
		return orders, err
	}

	fmt.Println("Success getting orders data")
	return orders, nil
}

func (r *repository) CreateOrder(order models.Orders) (models.Orders, error) {
	log.Printf("model: %+v", order)
	err := r.db.Create(&order).Error
	if err != nil {
		fmt.Println("Error creating order data:", err)
		return order, err
	}

	fmt.Println("Success creating order data")
	return order, nil
}

func (r *repository) UpdateOrder(order models.Orders) (models.Orders, error) {
	err := r.db.Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", order.ID).Updates(order).Error
	if err != nil {
		fmt.Println("Error updating order data:", err)
		return order, err
	}

	fmt.Println("Success updating order data")
	return order, nil
}

func (r *repository) DeleteOrder(orderID uint) error {
	order := models.Orders{}
	err := r.db.Where("id = ?", orderID).Delete(&order).Error
	if err != nil {
		fmt.Println("Error deleting order data:", err)
		return err
	}

	fmt.Println("Success deleting order data")
	return nil
}
