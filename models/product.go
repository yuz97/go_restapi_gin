package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	ProductName string `gorm:"type:varchar(30)" json:"product_name"`
	Description string `gorm:"type:text" json:"description"`
}
