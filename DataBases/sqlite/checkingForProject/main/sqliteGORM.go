package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	//"fmt"
)

type Contact struct {
	MainUser string `gorm:`
	ContactUser string
	IsBlocked int
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Create table for model `User`
	db.CreateTable(&Contact{})


}
