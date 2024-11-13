package models

import "gorm.io/gorm"

type MasterWarehouse struct {
	gorm.Model
	WhsName string `json:"whs_name" binding:"required" gorm:"type:varchar(255)"`
}

type InputMasterWarehouse struct {
	WhsName string `json:"whs_name" binding:"required" gorm:"type:varchar(255)"`
}
