package controller

import (
	"TES-SAMB-GO/databases"
	"TES-SAMB-GO/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListTransaksiPenerimaanBarangDetail(ctx *gin.Context) {
	var data []models.TransaksiPenerimaanBarangDetail

	if err := databases.DB.Table("transaksi_penerimaan_barang_details").Preload("Header").Preload("Header.MasterWarehouse").Preload("Header.MasterSupplier").Preload("Product").Find(&data).Error; err != nil {
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

func CreateTransaksiPenerimaanBarangDetail(ctx *gin.Context) {
	var input models.InputTransaksiPenerimaanBarangDetail

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	data := models.TransaksiPenerimaanBarangDetail{
		TrxInIDF:         input.TrxInIDF,
		TrxInDProductIdf: input.TrxInDProductIdf,
		TrxInDQtyDus:     input.TrxInDQtyDus,
		TrxInDQtyPcs:     input.TrxInDQtyPcs,
	}

	if err := databases.DB.Table("transaksi_penerimaan_barang_details").Create(&data).Error; err != nil {
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

func UpdateTransaksiPenerimaanBarangDetail(ctx *gin.Context) {
	var input models.InputTransaksiPenerimaanBarangDetail
	var data models.TransaksiPenerimaanBarangDetail

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	if err := databases.DB.Table("transaksi_penerimaan_barang_details").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
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

	data.TrxInIDF = input.TrxInIDF
	data.TrxInDProductIdf = input.TrxInDProductIdf
	data.TrxInDQtyDus = input.TrxInDQtyDus
	data.TrxInDQtyPcs = input.TrxInDQtyPcs

	if err := databases.DB.Table("transaksi_penerimaan_barang_details").Save(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "update data success",
		"data":    input,
	})
}

func DeleteTransaksiPenerimaanBarangDetail(ctx *gin.Context) {
	var data models.TransaksiPenerimaanBarangDetail

	if err := databases.DB.Table("transaksi_penerimaan_barang_details").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
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

	if err := databases.DB.Table("transaksi_penerimaan_barang_details").Delete(&data).Error; err != nil {
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
