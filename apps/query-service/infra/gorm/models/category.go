package models

type Category struct {
	InternalId int    `gorm:"column:c_key;primary_key"`
	Id         string `gorm:"column:c_id;unique"`
	Name       string `gorm:"column:c_name"`
}
