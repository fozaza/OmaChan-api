package table

import (
	"github.com/OmaChan/database"
	"github.com/OmaChan/module"
	"gorm.io/gorm"
)

// User is main data user
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Level    int
}

// UserInput for login
type UserInput struct {
	Name     string
	Email    string
	Password string
}

type UserLogin struct {
	Email    string
	Password string
}

type UserRetrun struct {
	Email      string
	Level_user int
}

// func Create user
func Cr_user(user UserInput) module.ErrorOmaChan {
	var new_user User
	new_user.Email = user.Email
	new_user.Name = user.Name
	new_user.Level = 1

	password, err := module.Cr_pw(user.Password)
	if err.Err != nil {
		return err
	}
	new_user.Password = password

	db := database.Get_db()
	db.Create(&new_user)
	return module.New_ErrorOmChan()
}

// func login user
func Login(user_input UserLogin) (UserRetrun, error) {
	var user_login UserRetrun
	var user User
	db := database.Get_db()
	result := db.Debug().Where("email=?", user_input.Email).First(&user)
	if result.Error != nil {
		return user_login, result.Error
	}

	if err := module.Ch_pw(user_input.Password, user.Password); err != nil {
		return user_login, nil
	}
	user_login.Email = user.Email
	user_login.Level_user = user.Level
	return user_login, nil
}
