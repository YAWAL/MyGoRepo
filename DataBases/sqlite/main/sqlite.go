package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db, _ = sql.Open("sqlite3", "DataBase/sqlite/client_db")

	createDBquery = "create table if not exists contacts (id primary key int, name varchar (20), birthday date, isblocked boolean)"
)

type Contact struct {
	Id        int
	Name      string
	Birtday   int
	IsBlocked bool
}

func createDB(query string) error {

	result, err := db.Exec(createDBquery)
	if err != nil {
		return err
	}

	fmt.Print(result)
	return nil
}

func (c *Contact) InsertIntoContact(Id int, Name string, Birtday int, IsBlocked bool) (bool, error) {

	tx, _ := db.Begin()
	stmp, _ := tx.Prepare("insert into contacts(id, name, birthday, isblocked) values (?, ?, ?)")

	_, err := stmp.Exec(c.Id, c.Birtday, c.IsBlocked)

	tx.Commit()

	return true, err

}

func main() {
	createDB(createDBquery)
	//c := Contact{}
	//c.InsertIntoContact(56, "56767", 1568904576, true)

}
