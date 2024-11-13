package models

import (
	"time"

	"gorm.io/gorm"
)

type TransaksiPenerimaanBarangHeader struct {
	gorm.Model
	TrxInNo         string          `json:"trx_in_no" binding:"required" gorm:"type:varchar(50);not null"`
	WhsIdf          uint            `json:"whs_idf" binding:"required" gorm:"not null"`
	MasterWarehouse MasterWarehouse `gorm:"foreignKey:WhsIdf;references:ID"`
	TrxInDate       time.Time       `json:"trx_in_date" binding:"required" gorm:"not null"`
	TrxInSuppIdf    uint            `json:"trx_in_supp_idf" binding:"required" gorm:"not null"`
	MasterSupplier  MasterSupplier  `gorm:"foreignKey:TrxInSuppIdf;references:ID"`
	TrxInNotes      string          `json:"trx_in_notes" binding:"required" gorm:"type:varchar(255);not null"`
}

type InputTransaksiPenerimaanBarangHeader struct {
	TrxInNo      string    `json:"trx_in_no" binding:"required" gorm:"type:varchar(50);not null"`
	WhsIdf       uint      `json:"whs_idf" binding:"required" gorm:"not null"`
	TrxInDate    time.Time `json:"trx_in_date" binding:"required" gorm:"not null"`
	TrxInSuppIdf uint      `json:"trx_in_supp_idf" binding:"required" gorm:"not null"`
	TrxInNotes   string    `json:"trx_in_notes" binding:"required" gorm:"type:varchar(255);not null"`
}
