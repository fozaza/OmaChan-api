package main

import (
	"os"

	"github.com/OmaChan/database"
	"github.com/OmaChan/install"
	"github.com/OmaChan/module/jp2a"
	"github.com/OmaChan/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
	if err := database.Open(os.Getenv("db_path")); err != nil {
		panic(err.Error())
	}
	install.Install_table()
	install.Install_root()
	jp2a.Print("~/Documents/golang/OmaChan/module/jp2a/image/Dragon_Comic.jpg")
	server.OpenServer()
}
