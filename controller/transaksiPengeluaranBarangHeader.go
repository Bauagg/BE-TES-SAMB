package controller

import (
	"TES-SAMB-GO/databases"
	"TES-SAMB-GO/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListTransaksiPengeluaranBarangHeader(ctx *gin.Context) {
	var data []models.TransaksiPengeluaranBarangHeader

	if err := databases.DB.Table("transaksi_pengeluaran_barang_headers").Preload("MasterWarehouse").Preload("MasterSupplier").Find(&data).Error; err != nil {
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

func CreateTransaksiPengeluaranBarangHeader(ctx *gin.Context) {
	var input models.InputTransaksiPengeluaranBarangHeader

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	data := models.TransaksiPengeluaranBarangHeader{
		TrxOutNo:      input.TrxOutNo,
		WhsIdf:        input.WhsIdf,
		TrxOutDate:    input.TrxOutDate,
		TrxOutSuppIdf: input.TrxOutSuppIdf,
		TrxOutNotes:   input.TrxOutNotes,
	}

	if err := databases.DB.Table("transaksi_pengeluaran_barang_headers").Create(&data).Error; err != nil {
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

func UpdateTransaksiPengeluaranBarangHeader(ctx *gin.Context) {
	var input models.InputTransaksiPengeluaranBarangHeader
	var data models.TransaksiPengeluaranBarangHeader

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	if err := databases.DB.Table("transaksi_pengeluaran_barang_headers").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
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

	data.TrxOutNo = input.TrxOutNo
	data.WhsIdf = input.WhsIdf
	data.TrxOutDate = input.TrxOutDate
	data.TrxOutSuppIdf = input.TrxOutSuppIdf
	data.TrxOutNotes = input.TrxOutNotes

	if err := databases.DB.Table("transaksi_pengeluaran_barang_headers").Save(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "update data terbaru",
		"data":    input,
	})
}

func DeleteTransaksiPengeluaranBarangHeader(ctx *gin.Context) {
	var data models.TransaksiPengeluaranBarangHeader

	if err := databases.DB.Table("transaksi_pengeluaran_barang_headers").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
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

	if err := databases.DB.Table("transaksi_pengeluaran_barang_headers").Delete(&data).Error; err != nil {
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
