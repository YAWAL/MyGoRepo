package main

import (
	"os"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native" // Native engine
	 _ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
	"fmt"
)


func openDB() mysql.Conn {
	address := "127.0.0.1:3306"
	user := "root"
	pass := "root"
	dbname := "users"
	db := mysql.New("tcp", "", address, user, pass, dbname)
	err := db.Connect()
	if err != nil {
		panic(err.Error())
	}
	return db
}


func main() {
	db := openDB()

	rows, _, err := db.Query("select * from contacts")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", db.Ping())


	for _, row := range rows {
		for _, col := range row {
			if col == nil {
				// col has NULL value
			} else {
				// Do something with text in col (type []byte)
			}
		}
		// You can get specific value from a row
		val1 := row[1].([]byte)

		// You can use it directly if conversion isn't needed
		os.Stdout.Write(val1)
		fmt.Print(val1)

		// You can get converted value
		//number := row.Int(0)      // Zero value
		//str    := row.Str(1)      // First value
		//bignum := row.MustUint(2) // Second value

		// You may get values by column name
		//first := res.Map("FirstColumn")
		//second := res.Map("SecondColumn")
		//SecondColumnval1, val2 := row.Int(first), row.Str(second)
	}

}