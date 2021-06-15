package user

import (
	"log"

	"github.com/fahruluzi/pos-mini/src/utils"
	"github.com/fahruluzi/pos-mini/src/utils/db"
	"gorm.io/gorm"
)

func Save(user Users) (insertedUUID string, result *gorm.DB) {
	db := db.GetDB()
	log.Print("user obj before repo save: ", user)
	result = db.Create(&user)
	insertedUUID = user.UUID
	return
}

func GetUserByEmail(email string) (Users, error) {
	db := db.GetDB()
	userLogin := Users{}

	result := db.First(&userLogin, "email = ?", email)
	if result.Error != nil {
		return Users{}, result.Error
	}
	return userLogin, nil
}

func GetUsers(pagination *utils.Pagination) ([]UsersList, error) {
	db := db.GetDB()
	var getUsers []UsersList

	offset := (pagination.Page - 1) * pagination.Limit

	if err := db.Model(&Users{}).Limit(pagination.Limit).Offset(offset).Find(&getUsers).Error; err != nil {
		return getUsers, err
	}

	return getUsers, nil
}

func CountUsers() (int, error) {
	db := db.GetDB()
	var usersCount int64

	if err := db.Model(&Users{}).Count(&usersCount).Error; err != nil {
		return 0, err
	}

	return int(usersCount), nil
}
