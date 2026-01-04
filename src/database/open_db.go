package database

import (
	"fmt"
	"os"
	//"github.com/OmaChan/module"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func Open(path string) error {
	fmt.Println("Open Server")
	//file := path
	//module.Get_file_path(path)
	// if err != nil {
	// 	return err
	// }

	hostLocal := "host="
	env := os.Getenv("OPEN_WITH_DOCKER")
	if env == "" {
		hostLocal += "localhost"
	} else {
		hostLocal += "postgres"
	}

	dsn := hostLocal + " user=root password=qqee22rr43 dbname=oma_chan_data port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
