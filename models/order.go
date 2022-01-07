package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	ProductRef int     `json:"product_id"`
	Product    Product `gorm:"foreignKey:ProductRef"`
	UserRef    int     `json:"user_id"`
	User       User    `gorm:"foreignKey:UserRef"`
}
