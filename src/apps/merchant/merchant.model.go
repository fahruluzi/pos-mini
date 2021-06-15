package merchant

import "github.com/fahruluzi/pos-mini/src/utils"

type Merchants struct {
	UUID   string `gorm:"type:varchar(36);primaryKey;not null" json:"uuid"`
	Name   string `gorm:"type:varchar(255);not null" json:"name"`
	Telp   string `gorm:"type:varchar(15);not null" json:"telp"`
	Jenis  string `gorm:"type:varchar(32);not null" json:"jenis"`
	Alamat string `gorm:"type:varchar(255);not null" json:"alamat"`
	utils.TimestampModel
}
