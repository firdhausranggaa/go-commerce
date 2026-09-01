package models

type Cart struct {
	ID        uint    `gorm:"primary_key" json:"id"`
	UserID    uint    `json:"user_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignkey:ProductID" json:"product"`
	Quantity  int     `json:"quantity"`
}