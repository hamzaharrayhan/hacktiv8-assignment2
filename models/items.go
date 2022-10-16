package models

type Items struct {
	ID          uint   `gorm:"primaryKey" json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint
}
