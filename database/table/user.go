package table

import (
	"errors"
	"fmt"
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

func (user User) To_retrun() UserGetRetrun {
	var user_retrun = UserGetRetrun{
		Name:  user.Name,
		Email: user.Email,
		Level: user.Level,
	}
	return user_retrun
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

type UserGetRetrun struct {
	Name  string
	Email string
	Level int
}

func Map_user(user []User) []UserGetRetrun {
	user_new := []UserGetRetrun{}
	for _, u := range user {
		user_new = append(user_new, u.To_retrun())
	}
	return user_new
}

// user for query user. user req name,email
type QueryUser struct {
	Email       string
	Name        string
	StartIntdex uint
	MaxOuput    uint
}

func (query QueryUser) DataMapQuery() string {
	var email = "email="
	if query.Email == "" {
		email = ""
	}

	var name = "name="
	if query.Name == "" {
		name = ""
	}

	return_query := fmt.Sprintf("%s %s %s %s", email, query.Email, name, query.Name)
	return return_query
}

type RemoveUserWithAdmin struct {
	Email []string
	Admin UserLogin
}

type Password string

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
	result := db.Debug().Where("emailอีเมล=?", user_login.Email).First(&user)
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

// get user with email name
func Gt_user(query QueryUser) (UserGetRetrun, error) {
	var user User
	db := database.Get_db()

	// map query
	var str_query = query.DataMapQuery()
	if result := db.Debug().Where(str_query).First(&user); result.Error != nil {
		return user.To_retrun(), result.Error
	}

	return user.To_retrun(), nil
}

// get user all with
func Gt_all_user(query QueryUser) ([]UserGetRetrun, error) {
	var user []User
	db := database.Get_db()

	var str_query = query.DataMapQuery()
	if result := db.Debug().
		Limit(int(query.MaxOuput)).
		Offset(int(query.StartIntdex)).
		Where(str_query).
		Find(&user); result.Error != nil {
		return Map_user(user), result.Error
	}

	return Map_user(user), nil
}
