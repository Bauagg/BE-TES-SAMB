package controller

import (
	"TES-SAMB-GO/databases"
	"TES-SAMB-GO/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListTransaksiPengeluaranbarangDetail(ctx *gin.Context) {
	var data []models.TransaksiPengeluaranBarangDetail

	if err := databases.DB.Table("transaksi_pengeluaran_barang_details").Preload("Header").Preload("Header.MasterWarehouse").Preload("Header.MasterSupplier").Preload("Product").Find(&data).Error; err != nil {
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

func CreateTransaksiPengeluaranbarangDetail(ctx *gin.Context) {
	var input models.InputTransaksiPengeluaranBarangDetail

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	data := models.TransaksiPengeluaranBarangDetail{
		TrxOutIDF:         input.TrxOutIDF,
		TrxOutDProductIdf: input.TrxOutDProductIdf,
		TrxOutDQtyDus:     input.TrxOutDQtyDus,
		TrxOutDQtyPcs:     input.TrxOutDQtyPcs,
	}

	if err := databases.DB.Table("transaksi_pengeluaran_barang_details").Create(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"error":   false,
		"message": "Create data success",
		"data":    input,
	})
}

func UpdateTransaksiPengeluaranbarangDetail(ctx *gin.Context) {
	var input models.InputTransaksiPengeluaranBarangDetail
	var data models.TransaksiPengeluaranBarangDetail

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	if err := databases.DB.Table("transaksi_pengeluaran_barang_details").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
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

	data.TrxOutIDF = input.TrxOutIDF
	data.TrxOutDProductIdf = input.TrxOutDProductIdf
	data.TrxOutDQtyDus = input.TrxOutDQtyDus
	data.TrxOutDQtyPcs = input.TrxOutDQtyPcs

	if err := databases.DB.Table("transaksi_pengeluaran_barang_details").Save(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "Update data success",
		"data":    input,
	})
}

func DeleteTransaksiPengeluaranbarangDetail(ctx *gin.Context) {
	var data models.TransaksiPengeluaranBarangDetail

	if err := databases.DB.Table("transaksi_pengeluaran_barang_details").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
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

	if err := databases.DB.Table("transaksi_pengeluaran_barang_details").Delete(&data).Error; err != nil {
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
