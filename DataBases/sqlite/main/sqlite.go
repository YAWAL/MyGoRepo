package main

import "database/sql"

var (
	db, _    = sql.Open("sqlite3", "DataBase/sqlite/client_db")
	createDB = "create table if not exists contacts (id primary key int, name varchar (20), birthday date, isblocked boolean)"
)
