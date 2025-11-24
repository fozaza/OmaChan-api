package database

import (
	"fmt"
	"github.com/OmaChan/module"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

func Open(path string) error {
	fmt.Println("Open Server")
	file, err := module.Get_file_path(path)
	if err != nil {
		return err
	}

	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		return err
	}
	database = db
	fmt.Println("Open Success")
	return nil
}
func Get_db() *gorm.DB {
	if database == nil {
		panic("DataBase is Unsafe, Error db is nill")
	}
	return database
}
