package merchant

import "github.com/fahruluzi/pos-mini/src/utils"

type Merchants struct {
	UUID    string `gorm:"type:varchar(36);primaryKey;not null" json:"uuid"`
	Name    string `gorm:"type:varchar(255);not null" json:"name"`
	Phone   string `gorm:"type:varchar(15);not null" json:"phone"`
	Type    string `gorm:"type:varchar(32);not null" json:"type"`
	Address string `gorm:"type:varchar(255);not null" json:"address"`
	utils.TimestampModel
}
