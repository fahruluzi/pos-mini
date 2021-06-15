package user

import (
	"log"

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
