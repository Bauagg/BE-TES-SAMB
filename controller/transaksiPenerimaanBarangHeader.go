package controller

import (
	"TES-SAMB-GO/databases"
	"TES-SAMB-GO/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListTransaksiPenerimaanBarangHeader(ctx *gin.Context) {
	var data []models.TransaksiPenerimaanBarangHeader

	if err := databases.DB.Table("transaksi_penerimaan_barang_headers").Preload("MasterWarehouse").Preload("MasterSupplier").Find(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "list data success",
		"data":    data,
	})
}

func CreateTransaksiPenerimaanBarangHeader(ctx *gin.Context) {
	var input models.InputTransaksiPenerimaanBarangHeader

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	data := models.TransaksiPenerimaanBarangHeader{
		TrxInNo:      input.TrxInNo,
		WhsIdf:       input.WhsIdf,
		TrxInDate:    input.TrxInDate,
		TrxInSuppIdf: input.TrxInSuppIdf,
		TrxInNotes:   input.TrxInNotes,
	}

	if err := databases.DB.Table("transaksi_penerimaan_barang_headers").Create(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"error":   false,
		"message": "create data success",
		"data":    input,
	})
}

func UpdateTransaksiPenerimaanBarangHeader(ctx *gin.Context) {
	var input models.InputTransaksiPenerimaanBarangHeader
	var data models.TransaksiPenerimaanBarangHeader

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	if err := databases.DB.Table("transaksi_penerimaan_barang_headers").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "data Not Found",
			})
			return
		}

		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	data.TrxInNo = input.TrxInNo
	data.WhsIdf = input.WhsIdf
	data.TrxInDate = input.TrxInDate
	data.TrxInSuppIdf = input.TrxInSuppIdf
	data.TrxInNotes = input.TrxInNotes

	if err := databases.DB.Table("transaksi_penerimaan_barang_headers").Save(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"error":   false,
		"message": "update data success",
		"data":    input,
	})
}

func DeleteTransaksiPenerimaanBarangHeader(ctx *gin.Context) {
	var data models.TransaksiPenerimaanBarangHeader

	if err := databases.DB.Table("transaksi_penerimaan_barang_headers").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "data Not Found",
			})
			return
		}

		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	if err := databases.DB.Table("transaksi_penerimaan_barang_headers").Delete(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "delete data success",
	})
}
