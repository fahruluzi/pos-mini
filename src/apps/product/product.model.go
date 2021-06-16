package product

import (
	"github.com/fahruluzi/pos-mini/src/apps/merchant"
	"github.com/fahruluzi/pos-mini/src/utils"
	"gorm.io/gorm"
)

type Products struct {
	UUID         string             `gorm:"type:varchar(36);primaryKey;not null" json:"uuid"`
	Name         string             `gorm:"type:varchar(255);not null" json:"name"`
	SKU          string             `gorm:"index:sku_index;type:varchar(255);not null;unique" json:"sku"`
	Price        int32              `gorm:"not null" json:"price"`
	ActualPrice  int32              `gorm:"not null" json:"actual_price"`
	Stock        int32              `gorm:"not null" json:"stock"`
	Image        string             `gorm:"type:varchar(255);not null" json:"image"`
	MerchantUuid string             `gorm:"type:varchar(36);not null" json:"merchant_uuid"`
	Merchant     merchant.Merchants `gorm:"->:false;<-:create;foreignKey:MerchantUuid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	utils.TimestampModel
}

func (p *Products) BeforeCreate(tx *gorm.DB) (err error) {
	p.UUID = utils.GenerateUuid()
	if !utils.ValidateUuid(p.UUID) {
		return err
	}

	return nil
}

type ListProducts struct {
	UUID         string `gorm:"type:varchar(36);primaryKey;not null" json:"uuid"`
	Name         string `gorm:"type:varchar(255);not null" json:"name"`
	SKU          string `gorm:"index:sku_index;type:varchar(255);not null;unique" json:"sku"`
	Price        int32  `gorm:"not null" json:"price"`
	ActualPrice  int32  `gorm:"not null" json:"actual_price"`
	Stock        int32  `gorm:"not null" json:"stock"`
	Image        string `gorm:"type:varchar(255);not null" json:"image"`
	MerchantUuid string `gorm:"type:varchar(36);not null" json:"merchant_uuid"`
	utils.TimestampModel
}
