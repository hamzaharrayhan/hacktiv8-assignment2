package payload

import (
	"assignment2/models"
	"time"
)

type GetOrderPayload struct {
	ID int `uri:"orderId" binding:"required"`
}

type OrderPayload struct {
	CustomerName string         `json:"customerName"  binding:"required"`
	OrderedAt    time.Time      `json:"orderedAt"  binding:"required"`
	Items        []models.Items `json:"items"  binding:"required"`
}
