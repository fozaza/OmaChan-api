package install

import (
	"fmt"

	"github.com/OmaChan/database"
	"github.com/OmaChan/database/table"
	"github.com/OmaChan/module"
)

func Install_root() {
	var root table.User
	db := database.Get_db()
	result := db.Debug().Where("name=?", "root").First(&root)
	if result.Error == nil {
		fmt.Println("OmaChan >>> OmaChan has Root")
		return
	}

	// input data
	var password string
	var email string
	fmt.Print("OmaChan >>> Create root input password: ")
	fmt.Scanln(&password)
	fmt.Print("OmaChan >>> Create root input email: ")
	fmt.Scanln(&email)

	// Create root
	// Gen password
	fmt.Println("Password " + password)
	password_hash, result_err := module.Cr_pw(password)
	if result_err.Err != nil {
		panic(result_err.Err)
	}

	// check password
	if err := module.Ch_pw(password, password_hash); err != nil {
		panic(err.Error())
	}

	// add root to database
	root.Name = "root"
	root.Password = password_hash
	root.Email = email
	root.Level = 5
	result = db.Create(&root)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println("OmaChan >>> Create root success")
}
