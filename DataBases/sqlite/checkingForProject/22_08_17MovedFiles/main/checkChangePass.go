package main

import (
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/jinzhu/gorm"

	"fmt"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/8tomat8/SSU-Golang-252-Chat/settingService"
)

//type Authentification struct {
//	UserName string `gorm:"primary_key"`
//	Password string
//	NickName string
//}

func ChangePass2(userName, newPass string) (bool, error) {
	var ayth  Authentification
	db, err := gorm.Open("sqlite3", "./database/messager.db")
	loger.Log.Println("conn established")

	defer db.Close()

	if err != nil {
		loger.Log.Errorf("DB error has occurred: ", err)
		return false, err
	}
	var result settingService.AuthResult //variable for storing result of querying into AuthResult struct
	// SELECT user_name, nick_name, password, birthday, about_user FROM users
	// WHERE main_user = "mainUser value from request body"
	// AND contact_user = "contactUser value from request body"
	db.Table("users").
		Select("user_name, nick_name, password").
		Where("user_name = ?", userName).
		Scan(&result)
	oldPass := result.Password
	if newPass == oldPass {
		loger.Log.Warnf(" Password is the same as old")
		return false, nil
	}
	fmt.Println(result)
	// UPDATE users SET password = "newPass value from request body"
	// WHERE user_name = "userName value from request header"
	db.Model(&messager).Where("user_name = ?", userName).Update("password", newPass)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}

func main() {
	boolVal, err := ChangePass2("", "")
	fmt.Println(boolVal, err)
}
