package main

import (
	"github.com/OmaChan/module/jp2a"
	"github.com/OmaChan/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
	jp2a.Print("~/Documents/golang/OmaChan/module/jp2a/image/Dragon_Comic.jpg")
	server.OpenServer()
}
