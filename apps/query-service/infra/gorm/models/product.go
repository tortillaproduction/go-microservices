package models

type Product struct {
	InternalId   int    `gorm:"column:p_key;primary_key"`
	Id           string `gorm:"column:p_id;unique"`
	Name         string `gorm:"column:p_name"`
	Price        uint32 `gorm:"column:p_price"`
	CategoryId   string `gorm:"column:c_id"`
	CategoryName string `gorm:"column:c_name"`
}
