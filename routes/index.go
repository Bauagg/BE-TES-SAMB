package routes

import (
	"TES-SAMB-GO/controller"

	"github.com/gin-gonic/gin"
)

func IndexRouter(app *gin.Engine) {
	router := app

	// Router Supplier
	router.GET("/api/supplier", controller.ListMasterSupplier)
	router.POST("/api/supplier", controller.CreateMasterSupplier)
	router.PUT("/api/supplier/:id", controller.UpdateMasterSupplier)
	router.DELETE("/api/supplier/:id", controller.DeleteMasterSupplier)

	// Router Product
	router.GET("/api/product", controller.ListMasterProduct)
	router.POST("/api/product", controller.CreateMasterProduct)
	router.PUT("/api/product/:id", controller.UpdateMasterProduct)
	router.DELETE("/api/product/:id", controller.DeleteMasterProduct)

	// Router Customer
	router.GET("/api/customer", controller.ListMasterCustomer)
	router.POST("/api/customer", controller.CreateMasterCustomer)
	router.PUT("/api/customer/:id", controller.UpdateMasterCustomer)
	router.DELETE("/api/customer/:id", controller.DeleteMasterCustomer)

	// Router Warehouse
	router.GET("/api/warehouse", controller.ListMasterWarehouse)
	router.POST("/api/warehouse", controller.CreateMasterWarehouse)
	router.PUT("/api/warehouse/:id", controller.UpdateMasterWarehouse)
	router.DELETE("/api/warehouse/:id", controller.DeleteMasterWarehouse)

	// Router Transaksi Penerimaan Barang Header
	router.GET("/api/transaksi-penerima-barang-header", controller.ListTransaksiPenerimaanBarangHeader)
	router.POST("/api/transaksi-penerima-barang-header", controller.CreateTransaksiPenerimaanBarangHeader)
	router.PUT("/api/transaksi-penerima-barang-header/:id", controller.UpdateTransaksiPenerimaanBarangHeader)
	router.DELETE("/api/transaksi-penerima-barang-header/:id", controller.DeleteTransaksiPenerimaanBarangHeader)

	// Router Transaksi Penerimaan Barang Detail
	router.GET("/api/transaksi-penerima-barang-detail", controller.ListTransaksiPenerimaanBarangDetail)
	router.POST("/api/transaksi-penerima-barang-detail", controller.CreateTransaksiPenerimaanBarangDetail)
	router.PUT("/api/transaksi-penerima-barang-detail/:id", controller.UpdateTransaksiPenerimaanBarangDetail)
	router.DELETE("/api/transaksi-penerima-barang-detail/:id", controller.DeleteTransaksiPenerimaanBarangDetail)

	// Router Transaksi Pengeluaran Barang Header
	router.GET("/api/transaksi-pengeluaran-barang-header", controller.ListTransaksiPengeluaranBarangHeader)
	router.POST("/api/transaksi-pengeluaran-barang-header", controller.CreateTransaksiPengeluaranBarangHeader)
	router.PUT("/api/transaksi-pengeluaran-barang-header/:id", controller.UpdateTransaksiPengeluaranBarangHeader)
	router.DELETE("/api/transaksi-pengeluaran-barang-header/:id", controller.DeleteTransaksiPengeluaranBarangHeader)

	// Router Transaksi Pengeluaran Barang Detail
	router.GET("/api/transaksi-pengeluaran-barang-detail", controller.ListTransaksiPengeluaranbarangDetail)
	router.POST("/api/transaksi-pengeluaran-barang-detail", controller.CreateTransaksiPengeluaranbarangDetail)
	router.PUT("/api/transaksi-pengeluaran-barang-detail/:id", controller.UpdateTransaksiPengeluaranbarangDetail)
	router.DELETE("/api/transaksi-pengeluaran-barang-detail/:id", controller.DeleteTransaksiPengeluaranbarangDetail)
}
