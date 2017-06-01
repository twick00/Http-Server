package command

import (
	"database/sql"
	"io/ioutil"
)

var db *sql.DB

func connectdb() *sql.DB {
	if db == nil {
		ini, err := ioutil.ReadFile("./pass.txt")
		if err != nil {
			panic(err)
		}
		inistr := string(ini)
		db, err = sql.Open("mysql", inistr)
		if err != nil {
			panic(err)
		}
	}
	return db
}

func readsqlscript() string {
	in, err := ioutil.ReadFile("./migrate.sql")
	if err != nil {
		panic(err)
	}
	s := string(in[:])
	return s
}

func main() {
	connectdb()
	out, err := db.Prepare(readsqlscript())
	if err != nil {
		panic(err)
	}
	_, err = out.Exec()
	if err != nil {
		panic(err)
	}
	db.Close()
}
