package table

import (
	"errors"
	"strings"

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
	Email string
	Level int
}

// func Create user
func Cr_user(user UserInput) module.ErrorOmaChan {
	var new_user User
	name := string(user.Name)
	strings.ToLower(name)
	strings.TrimSpace(name)

	email := user.Email
	strings.ToLower(email)
	strings.TrimSpace(email)

	new_user.Email = email
	new_user.Name = name
	new_user.Level = 1

	if new_user.Name == "root" {
		return module.New_ErrorOmChan().Errors("OmaChan >>> u can't use root")
	}

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
	user_login.Level = user.Level
	return user_login, nil
}

// Chage level
func Ch_le(admin UserRetrun, email string, new_level int) error {
	if new_level > 4 && new_level < 1 {
		return errors.New("OmaChan >>> u can't add new level >4 and <1")
	}

	var user User
	db := database.Get_db()
	result := db.Debug().Where("email=?", email).First(&user)
	if result.Error != nil {
		return result.Error
	}

	if user.Name == "root" {
		return errors.New("OmaChan >>> error u can't change root")
	}
	user.Level = new_level
	db.Save(&user)

	return nil
}

// Remove user (self)
func Rm_self(user_login UserLogin) error {
	var user UserLogin
	db := database.Get_db()
	result := db.Debug().Where("email=?", user_login.Email).First(&user)
	if result.Error != nil {
		return errors.New("OmaChan >>> not found email")
	}

	// check user
	if err := module.Ch_pw(user_login.Password, user.Password); err != nil {
		return err
	}

	result = db.Debug().Delete(user)
	if result.Error != nil {
		return errors.New("OmaChan >>> error delete id")
	}
	return nil
}

// Remove with admin or root
func Rm_user(id_admin UserLogin, user_email []string) (string, error) {
	// get admin or root id with email
	var admin UserLogin
	db := database.Get_db()
	result := db.Where(id_admin.Email).First(&admin)
	if result.Error != nil {
		return "", errors.New("OmaChan >>> not found email")
	}
	// admin or root password
	if err := module.Ch_pw(id_admin.Password, admin.Password); err != nil {
		return "", err
	}

	// delete user all with email
	var processs string // for msg loop delete user
	for _, email := range user_email {
		result = db.Debug().Where("email=?", email).Delete(&User{})
		if result.Error != nil {
			processs = email + ": Error\n"
			continue
		}
		processs = email + ": Success\n"
	}
	return processs, nil
}
