package main

import (
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	//"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"fmt"

	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func ChangeNickName2(userName string, nickName string, pass string) (bool, error) {

	db, err := gorm.Open("sqlite3", "../database/messager_test.db")
	if err == nil {
		loger.Log.Println("conn established")
	}
	fmt.Println(err)
	db.Debug()
	defer db.Close()

	authentification := messageService.Authentification{UserName: userName, NickName: nickName, Password: pass}
	db.Model(&authentification).Update("nick_name", "newnick")
	if err != nil {
		loger.Log.Errorf("DB error has occurred: ", err)
		return false, err
	}

	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}

type AuthResult struct {
	UserName string
	NickName string
	Password string
}

func read() {
	db, err := gorm.Open("sqlite3", "../database/messager_test.db")
	if err == nil {
		loger.Log.Println("conn established")
	}
	fmt.Println(err)
	db.Debug()
	defer db.Close()
	userName := "Yurii"
	var result AuthResult //variable for storing result of querying into AuthResult struct
	// SELECT user_name, nick_name, password, birthday, about_user FROM users
	// WHERE main_user = "mainUser value from request body"
	// AND contact_user = "contactUser value from request body"
	db.Table("authentifications").
		Select("user_name, nick_name, password").
		Where("user_name = ?", userName).
		Scan(&result)

	fmt.Println(result)
}

func main() {
	//ChangeNickName2("Yurii", "zzzz", "pass")
	//ok, err := ChangeNickName2("Yurii", "nick")
	//fmt.Println(ok, err)
	read()
}
