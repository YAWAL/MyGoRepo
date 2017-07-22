package main

import "database/sql"

var db, _ = sql.Open("sqlite3","DataBase/sqlite/client_db")
