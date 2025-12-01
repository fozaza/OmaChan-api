package table

import (
	"github.com/OmaChan/database"
	"gorm.io/gorm"
)

type HardWareData struct {
	gorm.Model
	Name         string `gorm:"uniqueIndex"`
	NameHardWare string
	Hardware     HardWare `gorm:"foreignKey:Name;references:NameHardWare;constraint:OnDelete:CASCADE"`
	Data         Data
}

type Data struct {
	Pm     float32
	Batter float32
}

func Cr_hwd(hardware HardWare, name string) error {
	db := database.Get_db()
	var hardware_data HardWareData
	hardware_data.Name = name
	hardware_data.Hardware = hardware
	if result := db.Debug().Create(&hardware); result.Error != nil {
		return result.Error
	}
	return nil
}

func Gt_hwd(name string) (HardWareData, error) {
	db := database.Get_db()
	var hardware_data HardWareData
	if result := db.Debug().Preload("HardWare").Where("name=?").First(&hardware_data); result.Error != nil {
		return hardware_data, result.Error
	}
	return hardware_data, nil
}
