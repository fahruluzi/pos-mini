package merchant

import (
	"log"

	"github.com/fahruluzi/pos-mini/src/utils/db"
	"gorm.io/gorm"
)

func Save(merchant Merchants) (insertedUUID string, result *gorm.DB) {
	db := db.GetDB()
	log.Print("merchant obj before repo save: ", merchant)
	result = db.Create(&merchant)
	insertedUUID = merchant.UUID
	return
}
