package main

import (
	"database/sql"
	"fmt"
)

var (
	db, _    = sql.Open("sqlite3", "DataBase/sqlite/client_db")
	createDBquery = "create table if not exists contacts (id primary key int, name varchar (20), birthday date, isblocked boolean)"
)

func createDB(query string) error {
	result, err := db.Exec(createDBquery)
	if err != nil {
		return err
	}
	fmt.Print(result)
	return nil
}