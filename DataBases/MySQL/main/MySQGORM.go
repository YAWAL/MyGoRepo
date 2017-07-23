package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql" // better driver to use with GORM
	"fmt"
)

type Contact struct {
	Id        int
	Name      string
	Birtday   int
	IsBlocked bool
}

type Employee struct {
	//gorm.Model
	FirstName string // Should Start with UpperCase !!!
	LastName  string // Should Start with UpperCase !!!
}

func Connect() *gorm.DB {
	db, err := gorm.Open(
		"mysql",
		"root:root@/users?charset=utf8&parseTime=True&loc=Local",
	)
	db.LogMode(true)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {

	db := Connect()

	err := Connect().DB().Ping()

	fmt.Println(err)

	db.CreateTable(Employee{})

	defer db.Close()

}
