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

func Update(user Users, uuid string) (updatedUUID string, result *gorm.DB) {
	db := db.GetDB()
	log.Print("user obj before repo update: ", user)
	result = db.Model(&Users{}).Where("uuid=?", uuid).Updates(&user)
	updatedUUID = uuid
	return
}

func Delete(uuid string) (deletedUUID string, result *gorm.DB) {
	db := db.GetDB()
	log.Print("user uuid before repo delete: ", uuid)
	result = db.Where("uuid = ?", uuid).Delete(&Users{})
	deletedUUID = uuid
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

func GetUser(uuid string) (UsersList, error) {
	db := db.GetDB()
	var getUser UsersList

	err := db.Model(&Users{}).Where("uuid = ?", uuid).First(&getUser).Error

	if err != nil {
		return UsersList{}, err
	}

	return getUser, nil
}

func GetUserWithPassword(uuid string) (Users, error) {
	db := db.GetDB()
	var getUser Users

	err := db.Model(&Users{}).Where("uuid = ?", uuid).First(&getUser).Error

	if err != nil {
		return Users{}, err
	}

	return getUser, nil
}
