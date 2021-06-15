package main

import (
	"log"
	"os"

	"github.com/fahruluzi/pos-mini/src/apps/merchant"
	"github.com/fahruluzi/pos-mini/src/apps/product"
	"github.com/fahruluzi/pos-mini/src/apps/user"
	"github.com/fahruluzi/pos-mini/src/utils/db"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&merchant.Merchants{})
	db.AutoMigrate(&user.Users{})
	db.AutoMigrate(&product.Products{})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}

	db := db.Init()
	Migrate(db)
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := Router()
	r.Run(":" + os.Getenv("PORT"))
}
