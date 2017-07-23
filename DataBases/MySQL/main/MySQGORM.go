package main

import (
	"github.com/jinzhu/gorm"
	//"github.com/ziutek/mymysql/mysql"
	//_ "github.com/ziutek/mymysql/native"  // Native engine
	//_ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
	"fmt"
)


var (
	address = "127.0.0.1:3306"
	user = "root"
	pass = "root"
	dbname = "users"
)


func main() {


	db, err := gorm.Open("MySQL","tcp", "", address, user, pass, dbname )

	if err != nil {
		panic(err.Error())
	}

	err = db.DB().Ping()

	fmt.Println(err)

	db.Close()
}
