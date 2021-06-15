package product

import (
	"github.com/fahruluzi/pos-mini/src/apps/merchant"
	"github.com/fahruluzi/pos-mini/src/utils"
)

type Products struct {
	UUID         string             `gorm:"type:varchar(36);primaryKey;not null" json:"uuid"`
	Name         string             `gorm:"type:varchar(255);not null" json:"name"`
	SKU          string             `gorm:"index:sku_index;type:varchar(255);not null;unique" json:"sku"`
	Price        string             `gorm:"type:int(16);not null" json:"price"`
	ActualPrice  string             `gorm:"type:int(16);not null" json:"actual_price"`
	Stock        string             `gorm:"type:int(16);not null" json:"stock"`
	Image        string             `gorm:"type:varchar(255);not null" json:"image"`
	MerchantUuid string             `gorm:"type:varchar(36);not null" json:"merchant_uuid"`
	Merchant     merchant.Merchants `gorm:"foreignKey:MerchantUuid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	utils.TimestampModel
}
