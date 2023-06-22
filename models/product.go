package models

type Product struct {
	Id          int64   `json:"id" gorm:"primary_key"`
	ProductName string  `gorm:"type:varchar(255)" json:"product_name"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"type:decimal(10,2)" json:"price"`
}
