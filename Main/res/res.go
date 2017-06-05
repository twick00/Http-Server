package dbconnect

import (
	"database/sql"
	"io/ioutil"

	//github.com/go-sql-driver/mysql is for connecting to MySQL server
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initconnectdb2() {
	var err error
	pass, err := ioutil.ReadFile("./res/pass.txt.")
	strpass := string(pass)
	db, err = sql.Open("mysql", strpass)
	if err != nil {
		panic(err)
	}
	//}
}

func readsqlscript() (string, string) {
	tabledata, err := ioutil.ReadFile("./res/migrate.sql")
	if err != nil {
		panic(err)
	}
	tab := string(tabledata[:])
	database, err := ioutil.ReadFile("./res/createdb.sql")
	if err != nil {
		panic(err)
	}
	dat := string(database[:])
	return dat, tab
}
func migrate(data string) {
	out, err := db.Prepare(data)
	if err != nil {
		panic(err)
	}
	_, err = out.Exec()
	if err != nil {
		panic(err)
	}
}

//Command connects to mysql and adds the necessary files for the server, then disconnects
func Command() *sql.DB {

	initconnectdb2()
	dat, tab := readsqlscript()
	migrate(dat)
	migrate(tab)
	// err := db.Close()
	// if err != nil {
	// 	panic(err)
	// }
	return db
}
