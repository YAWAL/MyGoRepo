package main

import (
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/jinzhu/gorm"
	"fmt"
)

func ChangeNickName2(userName, nickName string) (bool, error) {

	db, err := gorm.Open("sqlite3", "./DataBase/sqlite/messager_test.db")

	authentification := messageService.Authentification{UserName:userName}
	// UPDATE users SET nick_name = "nickName value from request body"
	// WHERE user_name = "userName value from request header"
	db.Model(&authentification).UpdateColumn("nick_name", nickName)
	if db.Error != nil {
		return false, err
	}
	return true, nil
}

func main() {
	ok, err := ChangeNickName2("Yuriy", "rome")
	fmt.Println(ok, err)
}
