package table

import (
	"github.com/OmaChan/database"
	"gorm.io/gorm"
)

type HardWare struct {
	gorm.Model
	Name  string `gorm:"uniqueIndex"`
	Title string
}

func (hardware HardWare) MapInput() HardWareInput {
	return HardWareInput{name: hardware.Name, title: hardware.Title}
}

type HardWareInput struct {
	name  string
	title string
}

func (hardware HardWareInput) MapHard() HardWare {
	return HardWare{Name: hardware.name, Title: hardware.title}
}

func Cr_ha(hardware_input HardWareInput) error {
	db := database.Get_db()
	var hardware = hardware_input.MapHard()
	if result := db.Debug().Create(&hardware); result.Error != nil {
		return result.Error
	}
	return nil
}

func De_ha(name string) error {
	db := database.Get_db()
	var hardware HardWare
	if result := db.Debug().Where("name=?", name).Delete(&hardware); result.Error != nil {
		return result.Error
	}
	return nil
}

func Ge_ha(name string) (HardWareInput, error) {
	db := database.Get_db()
	var hardware HardWare
	var hardware_return HardWareInput
	if result := db.Debug().Where("name=?", name); result.Error != nil {
		return hardware_return, result.Error
	}

	hardware_return = hardware.MapInput()
	return hardware_return, nil
}
func Ga_ha() {}
