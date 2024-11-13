package models

import "gorm.io/gorm"

type MasterProduct struct {
	gorm.Model
	ProductName string `json:"product_name" binding:"required" gorm:"type:varchar(255)"`
}

type InputMasterProduct struct {
	ProductName string `json:"product_name" binding:"required" gorm:"type:varchar(255)"`
}
