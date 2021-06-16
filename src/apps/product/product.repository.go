package product

import (
	"log"

	"github.com/fahruluzi/pos-mini/src/utils"
	"github.com/fahruluzi/pos-mini/src/utils/db"
	"gorm.io/gorm"
)

func Save(product Products) (insertedUUID string, result *gorm.DB) {
	db := db.GetDB()
	log.Print("product obj before repo save: ", product)
	result = db.Create(&product)
	insertedUUID = product.UUID
	return
}

func Update(product map[string]interface{}, uuid string) (updatedUUID string, result *gorm.DB) {
	db := db.GetDB()
	log.Print("product obj before repo update: ", product)
	result = db.Model(&Products{}).Where("uuid=?", uuid).Updates(product)
	updatedUUID = uuid
	return
}

func Delete(uuid string) (deletedUUID string, result *gorm.DB) {
	db := db.GetDB()
	log.Print("product uuid before repo delete: ", uuid)
	result = db.Where("uuid = ?", uuid).Delete(&Products{})
	deletedUUID = uuid
	return
}

func GetProductByEmail(email string) (Products, error) {
	db := db.GetDB()
	productLogin := Products{}

	result := db.First(&productLogin, "email = ?", email)
	if result.Error != nil {
		return Products{}, result.Error
	}
	return productLogin, nil
}

func GetProducts(pagination *utils.Pagination) ([]ListProducts, error) {
	db := db.GetDB()
	var getProducts []ListProducts

	offset := (pagination.Page - 1) * pagination.Limit

	if err := db.Model(&Products{}).Limit(pagination.Limit).Offset(offset).Find(&getProducts).Error; err != nil {
		return getProducts, err
	}

	return getProducts, nil
}

func CountProducts() (int, error) {
	db := db.GetDB()
	var productsCount int64

	if err := db.Model(&Products{}).Count(&productsCount).Error; err != nil {
		return 0, err
	}

	return int(productsCount), nil
}

func GetProduct(uuid string) (ListProducts, error) {
	db := db.GetDB()
	var getProduct ListProducts

	err := db.Model(&Products{}).Where("uuid = ?", uuid).First(&getProduct).Error

	if err != nil {
		return ListProducts{}, err
	}

	return getProduct, nil
}

func GetProductWithPassword(uuid string) (Products, error) {
	db := db.GetDB()
	var getProduct Products

	err := db.Model(&Products{}).Where("uuid = ?", uuid).First(&getProduct).Error

	if err != nil {
		return Products{}, err
	}

	return getProduct, nil
}
