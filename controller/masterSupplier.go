package controller

import (
	"TES-SAMB-GO/databases"
	"TES-SAMB-GO/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListMasterSupplier(ctx *gin.Context) {
	var data []models.MasterSupplier

	if err := databases.DB.Table("master_suppliers").Find(&data).Error; err != nil {
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

func CreateMasterSupplier(ctx *gin.Context) {
	var data models.MasterSupplier

	if errInput := ctx.ShouldBind(&data); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	if err := databases.DB.Table("master_suppliers").Create(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"error":   false,
		"message": "create data success",
		"data":    data,
	})
}

func UpdateMasterSupplier(ctx *gin.Context) {
	var input models.InputMasterSupplier
	var data models.MasterSupplier

	if errInput := ctx.ShouldBind(&input); errInput != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": errInput.Error(),
		})
		return
	}

	if err := databases.DB.Table("master_suppliers").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "data Master Suppliers Not Found",
			})
			return
		}

		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	data.SupplierName = input.SupplierName
	if err := databases.DB.Table("master_suppliers").Save(&data).Error; err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "update data success",
		"data":    data,
	})
}

func DeleteMasterSupplier(ctx *gin.Context) {
	var data models.MasterSupplier
	if err := databases.DB.Table("master_suppliers").Where("id = ?", ctx.Param("id")).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(404, gin.H{
				"error":   true,
				"message": "data Master Suppliers Not Found",
			})
			return
		}

		ctx.JSON(500, gin.H{
			"error":   true,
			"message": "Internal server error",
		})
		return
	}

	if err := databases.DB.Table("master_suppliers").Delete(&data).Error; err != nil {
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
