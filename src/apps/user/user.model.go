package user

import (
	"github.com/fahruluzi/pos-mini/src/apps/merchant"
	"github.com/fahruluzi/pos-mini/src/utils"
)

type Users struct {
	UUID         string             `gorm:"type:varchar(36);primaryKey;not null" json:"uuid"`
	Name         string             `gorm:"type:varchar(255);not null" json:"name"`
	Email        string             `gorm:"index:email_index;type:varchar(255);not null;unique" json:"email"`
	Password     string             `gorm:"type:varchar(255);not null" json:"password"`
	MerchantUuid string             `gorm:"type:varchar(36);not null" json:"merchant_uuid"`
	Merchant     merchant.Merchants `gorm:"foreignKey:MerchantUuid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	utils.TimestampModel
}
