package models

import "gorm.io/gorm"

type MasterSupplier struct {
	gorm.Model
	SupplierName string `json:"supplier_name" binding:"required" gorm:"type:varchar(255)"`
}

type InputMasterSupplier struct {
	SupplierName string `json:"supplier_name" binding:"required" gorm:"type:varchar(255)"`
}
