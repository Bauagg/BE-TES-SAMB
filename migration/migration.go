package migration

import (
	"TES-SAMB-GO/databases"
	"TES-SAMB-GO/models"
)

func Migration() {
	db := databases.GetDB()
	err := db.AutoMigrate(
		models.MasterSupplier{},
		models.MasterCustomer{},
		models.MasterProduct{},
		models.MasterWarehouse{},
		models.TransaksiPenerimaanBarangHeader{},
		models.TransaksiPenerimaanBarangDetail{},
		models.TransaksiPengeluaranBarangHeader{},
		models.TransaksiPengeluaranBarangDetail{},
	)

	if err != nil {
		panic("Failed to migrate: " + err.Error())
	}
}
