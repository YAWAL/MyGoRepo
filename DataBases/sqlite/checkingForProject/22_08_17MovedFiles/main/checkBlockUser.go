package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/jinzhu/gorm"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
)


func BlockUnblockUser2(mainUser string, contactUser string, isBlocked int) (bool, error) {

	db, err := gorm.Open("sqlite3", "./database/messager_test.db")
	if err == nil{
		loger.Log.Println("conn established")

	}
defer db.Close()
	contact := messageService.Contact{MainUser: mainUser, ContactUser: contactUser, IsBlocked: isBlocked}


	db.Model(&contact).Update("is_blocked", isBlocked)

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

func main() {
	BlockUnblockUser2("Yuriy", "Valeriy", 5)

}
