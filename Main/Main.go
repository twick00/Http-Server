package main

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"unicode/utf8"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

type mysql struct {
	Database *sql.DB
	Name     string
	Pass     string
}
type loginError struct {
	Error   error
	Message string
	Code    int
}

var login = mysql{}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

}

func getlogin(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func postlogin(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func stringchecker(scheck ...string) (int, error) {
	for _, str := range scheck {
		if str == "" {
			return -1, errors.New("Field is empty")
		}
		if utf8.RuneCountInString(str) > 32 {
			return -2, errors.New("Exceeds 32 string max")
		}
	}
	return 1, nil
}

func getroot(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var data string
	t, err := template.ParseFiles("./res/content.html")
	check(err)
	s1 := t.Lookup("content.html")
	s1.ExecuteTemplate(rw, "postgamedata", data)
}
func errorpage(rw http.ResponseWriter, r *http.Request, _ httprouter.Params, code int) {

}
func postroot(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	login.Name = r.FormValue("name")
	login.Pass = r.FormValue("pass")
	code, err := stringchecker(login.Pass, login.Name)

	if code != 1 {
		fmt.Println(err)
		errorpage(rw, r, nil, code)
		return
	}
	var db *sql.DB
	Connectdb(login.Name, login.Pass, "mysql", db)
	//TYLER ADD A PAGE FOR DATA #######################################

}
func routering() {
	router := httprouter.New()
	router.GET("/", getroot)
	router.POST("/", postroot)
	router.GET("/login", getlogin)
	router.POST("/login", postlogin)
}

//Connectdb establishes a connection to db where none exists.
func Connectdb(name string, pass string, dbtype string, db *sql.DB) *sql.DB {
	c := func() *sql.DB {
		if db == nil {
			db, err := sql.Open(dbtype, (name + "/" + pass + "@/"))
			check(err)
			return db
		}
		return db
	}
	return c()
}
