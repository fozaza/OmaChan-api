package table

import (
	"errors"

	"github.com/OmaChan/database"
	"gorm.io/gorm"
)

type OldHardWare struct {
	gorm.Model
	Name  string `gorm:"uniqueIndex"`
	Title string
}

func (hardware OldHardWare) MapInput() HardWareInput {
	return HardWareInput{name: hardware.Name, title: hardware.Title}
}

type HardWareInput struct {
	name  string
	title string
}

func (hardware HardWareInput) MapHard() OldHardWare {
	return OldHardWare{Name: hardware.name, Title: hardware.title}
}

type HardWare struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
	//NameHardWare string
	Title  string
	Pm     float32
	Batter float32
	Enable bool
}

type Data struct {
	Pm     float32
	Batter float32
}

func Cr_ha(hardwareNew HardWare) error {
	db := database.Get_db()
	var verfy HardWare
	if result := db.Debug().First(&verfy); result.Error == nil {
		return errors.New("Error found hardware in database")
	}
	if result := db.Debug().Create(&hardwareNew); result.Error != nil {
		return result.Error
	}
	return nil
}

func Up_data(new_data Data, name string) error {
	db := database.Get_db()
	var hardware HardWare
	if result := db.Debug().Where("name=?", name).First(&hardware); result.Error != nil {
		return errors.New("Error found hardware in database")
	}

	hardware.Pm = new_data.Pm
	hardware.Batter = new_data.Batter
	if result := db.Debug().Where("name=?", name).Save(&hardware); result.Error != nil {
		return result.Error
	}
	return nil
}

func Del_hrw(name string) error {
	db := database.Get_db()
	var hardware HardWare
	if result := db.Debug().Where("name=?", name).Unscoped().Delete(&hardware); result.Error != nil {
		return result.Error
	}
	return nil
}

// func Cr_ha(hardware_input HardWareInput) error {
// 	db := database.Get_db()
// 	var hardware = hardware_input.MapHard()
// 	if result := db.Debug().Create(&hardware); result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }
//
// func De_ha(name string) error {
// 	db := database.Get_db()
// 	var hardware HardWare
// 	if result := db.Debug().Where("name=?", name).Delete(&hardware); result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }
//
// func Ge_ha(name string) (HardWareInput, error) {
// 	db := database.Get_db()
// 	var hardware HardWare
// 	var hardware_return HardWareInput
// 	if result := db.Debug().Where("name=?", name); result.Error != nil {
// 		return hardware_return, result.Error
// 	}
//
// 	hardware_return = hardware.MapInput()
// 	return hardware_return, nil
// }
// func Ga_ha() {}
