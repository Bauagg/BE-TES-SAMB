package models

import (
	"time"

	"gorm.io/gorm"
)

type TransaksiPengeluaranBarangHeader struct {
	gorm.Model
	TrxOutNo        string          `json:"trx_out_no" binding:"required" gorm:"type:varchar(50);not null"`
	WhsIdf          uint            `json:"whs_idf" binding:"required" gorm:"not null"`
	MasterWarehouse MasterWarehouse `gorm:"foreignKey:WhsIdf;references:ID"`
	TrxOutDate      time.Time       `json:"trx_out_date" binding:"required" gorm:"not null"`
	TrxOutSuppIdf   uint            `json:"trx_out_supp_idf" binding:"required" gorm:"not null"`
	MasterSupplier  MasterSupplier  `gorm:"foreignKey:TrxOutSuppIdf;references:ID"`
	TrxOutNotes     string          `json:"trx_out_notes" binding:"required" gorm:"type:varchar(255);not null"`
}

type InputTransaksiPengeluaranBarangHeader struct {
	TrxOutNo      string    `json:"trx_out_no" binding:"required" gorm:"type:varchar(50);not null"`
	WhsIdf        uint      `json:"whs_idf" binding:"required" gorm:"not null"`
	TrxOutDate    time.Time `json:"trx_out_date" binding:"required" gorm:"not null"`
	TrxOutSuppIdf uint      `json:"trx_out_supp_idf" binding:"required" gorm:"not null"`
	TrxOutNotes   string    `json:"trx_out_notes" binding:"required" gorm:"type:varchar(255);not null"`
}
