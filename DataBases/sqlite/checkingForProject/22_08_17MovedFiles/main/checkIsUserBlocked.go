package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/8tomat8/SSU-Golang-252-Chat/settingService"
	"github.com/jinzhu/gorm"
)

type Result struct {
	MainUser    string
	ContactUser string
	IsBlocked   int
}

func IsUserBlocked2(mainUser, contactUser string) (isBlocked bool, err error) {
	isBlocked = true

	var result Result //variable for storing result of querying into ContactResult struct

	db, err := gorm.Open("sqlite3", "./database/messager_test.db")
	//if err != nil {
	//	loger.Log.Errorf("DB error has occurred: ", err)
	//	return true, err
	//}
	contact := settingService.Contact{MainUser: mainUser, ContactUser: contactUser}
	db.Where(&contact).First(&contact).Scan(result)
defer db.Close()
	// SELECT main_user, contact_user, is_blocked FROM contacts
	// WHERE main_user = "mainUser value from request body"
	// AND contact_user = "contactUser value from request body"
	//db.Table("contacts").
	//	Select("main_user, contact_user, is_blocked").
	//	Where("main_user = ? AND contact_user = ?", mainUser, contactUser).
	//	Scan(&result)
	//if db.Error != nil {
	//	err := errors.New("Bad parsing")
	//	loger.Log.Errorf("Error has occurred: ", err)
	//	return true, err
	//}
	switch result.IsBlocked {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return true, nil
	}
	return true, nil
}

func main() {
	isBlocked, err := IsUserBlocked2("Yuriy", "Valeriy")
	fmt.Println(isBlocked)
	fmt.Println(err)
}
