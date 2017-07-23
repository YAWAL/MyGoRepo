package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"  // better driver to use with GORM
	"fmt"
)

var (
	address = "127.0.0.1:3306"
	user    = "root"
	pass    = "root"
	dbname  = "users"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(
		"mysql",
		"root:root@/users?charset=utf8&parseTime=True&loc=Local",
	)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root@/users?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func main() {

	Connect()
	//connectDB()

	//err := connectDB().DB().Ping()
	err := Connect().DB().Ping()

	fmt.Println(err)

	//db.Close()
}
