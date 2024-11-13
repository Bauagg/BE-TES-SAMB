package models

import "gorm.io/gorm"

type TransaksiPengeluaranBarangDetail struct {
	gorm.Model
	TrxOutIDF         uint                             `json:"trx_out_idf" binding:"required" gorm:"not null"`
	Header            TransaksiPengeluaranBarangHeader `gorm:"foreignKey:TrxOutIDF;references:ID"`
	TrxOutDProductIdf uint                             `json:"trx_out_d_product_idf" binding:"required" gorm:"not null"`
	Product           MasterProduct                    `gorm:"foreignKey:TrxOutDProductIdf;references:ID"`
	TrxOutDQtyDus     int                              `json:"trx_out_d_qty_dus" binding:"required" gorm:"not null;default:0"`
	TrxOutDQtyPcs     int                              `json:"trx_out_d_qty_pcs" binding:"required" gorm:"not null;default:0"`
}

type InputTransaksiPengeluaranBarangDetail struct {
	TrxOutIDF         uint `json:"trx_out_idf" binding:"required" gorm:"not null"`
	TrxOutDProductIdf uint `json:"trx_out_d_product_idf" binding:"required" gorm:"not null"`
	TrxOutDQtyDus     int  `json:"trx_out_d_qty_dus" binding:"required" gorm:"not null;default:0"`
	TrxOutDQtyPcs     int  `json:"trx_out_d_qty_pcs" binding:"required" gorm:"not null;default:0"`
}
