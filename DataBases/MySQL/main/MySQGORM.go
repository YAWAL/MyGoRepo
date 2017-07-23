package main

import (

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql" // better driver to use with GORM
	"fmt"
	"log"
)

var (
	id int
	name string
)


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
	//db.CreateTable(Contact{1, "vasyl", 213456, true})
	//db.Exec("alter table add (id, name, birthday, isblocked) values ("2", "petro", 65556867, false)")

	num := db.Select("Select id, name from contacts")

	fmt.Printf("%v",num)
	defer db.Close()

	rows := db.Select("select id, name from users where id = 1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	//for rows.Find {
	//	//err := rows.Scan(&id, &name)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	log.Println(id, name)
	//}
	//err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
