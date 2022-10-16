package models

import "time"

type Orders struct {
	ID           uint `gorm:"primaryKey"`
	CustomerName string
	OrderedAt    time.Time
	Items        []Items `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
