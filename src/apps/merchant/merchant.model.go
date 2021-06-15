package merchant

import (
	"github.com/fahruluzi/pos-mini/src/utils"
	"gorm.io/gorm"
)

type Merchants struct {
	UUID    string `gorm:"type:varchar(36);primaryKey;not null" json:"uuid"`
	Name    string `gorm:"type:varchar(255);not null" json:"name"`
	Phone   string `gorm:"type:varchar(15);not null" json:"phone"`
	Address string `gorm:"type:varchar(255);not null" json:"address"`
	utils.TimestampModel
}

func (m *Merchants) BeforeCreate(tx *gorm.DB) (err error) {
	m.UUID = utils.GenerateUuid()
	if !utils.ValidateUuid(m.UUID) {
		return err
	}

	return nil
}
