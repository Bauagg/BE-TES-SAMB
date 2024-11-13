package models

import "gorm.io/gorm"

type TransaksiPenerimaanBarangDetail struct {
	gorm.Model
	TrxInIDF         uint                            `json:"trx_in_idf" binding:"required" gorm:"not null"`
	Header           TransaksiPenerimaanBarangHeader `gorm:"foreignKey:TrxInIDF;references:ID"`
	TrxInDProductIdf uint                            `json:"trx_in_d_product_idf" binding:"required" gorm:"not null"`
	Product          MasterProduct                   `gorm:"foreignKey:TrxInDProductIdf;references:ID"`
	TrxInDQtyDus     int                             `json:"trx_in_d_qty_dus" binding:"required" gorm:"not null;default:0"`
	TrxInDQtyPcs     int                             `json:"trx_in_d_qty_pcs" binding:"required" gorm:"not null;default:0"`
}

type InputTransaksiPenerimaanBarangDetail struct {
	TrxInIDF         uint `json:"trx_in_idf" binding:"required" gorm:"not null"`
	TrxInDProductIdf uint `json:"trx_in_d_product_idf" binding:"required" gorm:"not null"`
	TrxInDQtyDus     int  `json:"trx_in_d_qty_dus" binding:"required" gorm:"not null;default:0"`
	TrxInDQtyPcs     int  `json:"trx_in_d_qty_pcs" binding:"required" gorm:"not null;default:0"`
}
