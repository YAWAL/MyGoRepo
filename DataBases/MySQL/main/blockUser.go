package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql" // better driver to use with GORM
)

// discuss with team about db
// discuss with team about gorm

// should use structure from Svitlana's package
type Contact struct {
	Id        int
	Name      string
	Birtday   int
	IsBlocked bool
}

// IT'S WORKING !!!
// BlockUserByID function changes bool values IsBlocked in database to true
func BlockUserByID(id string) (bool, error) {
	// 1) establish db connection
	// 1.a) handle error
	db, err := gorm.Open(
		"mysql",
		"root:root@/users?charset=utf8&parseTime=True&loc=Local",
	)
	db.LogMode(true)
	if err != nil {
		panic(err.Error())
	}

	// 2) execyte query
	// 2.a) handle error
	query := "UPDATE contacts SET is_blocked = 1 WHERE id = " + id
	db.Exec(query)

	// 3)close db connection
	defer db.Close()

	return true, nil
}

// IT'S WORKING !!!
// BlockUserByName function changes bool values IsBlocked in database to true
func BlockUserByName(contactName string) (bool, error) {
	// 1) establish db connection
	// 1.a) handle error
	db, err := gorm.Open(
		"mysql",
		"root:root@/users?charset=utf8&parseTime=True&loc=Local",
	)
	db.LogMode(true)
	if err != nil {
		panic(err.Error())
	}

	// 2) execyte query
	// 2.a) handle error
	query := "UPDATE contacts SET is_blocked = 1 WHERE name = " + "\"" + contactName + "\""
	db.Exec(query)

	// 3)close db connection
	defer db.Close()

	return true, nil
}

// UnBlockUser function changes bool values IsBlocked in database to false
func (c *Contact) UnBlockUser(contactName string) (bool, error) {

	// 1) establish db connection
	// 1.a) handle error
	db, err := gorm.Open(
		"mysql",
		"root:root@/users?charset=utf8&parseTime=True&loc=Local",
	)
	db.LogMode(true)
	if err != nil {
		panic(err.Error())
	}

	// 2) execyte query
	// 2.a) handle error
	query := "UPDATE contacts SET is_blocked = 0 WHERE name = " + contactName
	db.Exec(query)

	// 3)close db connection
	defer db.Close()

	return true, nil
}

func main()  {
	//c := Contact{}
	//BlockUserByID("1")
	BlockUserByName("vasul")
}
