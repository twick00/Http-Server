package main

import (
	"database/sql"
	"testing"
)

//TestConnectdb tests Connectdb in Main.go
func TestConnectdb(t *testing.T) {
	name := "root"
	pass := "password"
	dbtype := "mysql"
	var db *sql.DB
	db = Connectdb(name, pass, dbtype, db)
	err := db.Ping()
	if err != nil {
		t.Error("Testing is unable to connect to the database, heres why", err)
	}
}
