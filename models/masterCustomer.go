package models

import "gorm.io/gorm"

type MasterCustomer struct {
	gorm.Model
	CustomerName string `json:"customer_name" binding:"required" gorm:"type:varchar(255)"`
}

type InputMasterCustomer struct {
	CustomerName string `json:"customer_name" binding:"required" gorm:"type:varchar(255)"`
}
