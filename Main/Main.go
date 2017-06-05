package main

import (
	"database/sql"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/sfreiberg/gotwilio"
	_ "github.com/sfreiberg/gotwilio"
)

var o sync.Once
var db *sql.DB
var data = struct {
	Title   string
	Year    int
	Genre   string
	Barcode int
	ID      int
}{}

var twilAuth = struct {
	sid   string
	token string
}{}

func check(e error) { //Simple error passing
	if e != nil {
		panic(e)
	}
}

func gethome(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles("./res/content.html")
	check(err)
	s1 := t.Lookup("content.html")
	s1.ExecuteTemplate(rw, "postgamedata", data)
}
func getroot(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles("./res/content.html")
	check(err)
	s1 := t.Lookup("content.html")
	s1.ExecuteTemplate(rw, "title", data)
}

func stringchecker(check string) string {
	if check == "" {
		return "0"
	}
	return check
}

func posthome(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	formTitle := r.FormValue("title")
	formYear := r.FormValue("year")
	formGenre := r.FormValue("genre")
	formBarcode := r.FormValue("barcode")
	formID := r.FormValue("ID")
	arrform := []string{formTitle, formYear, formGenre, formBarcode, formID}
	for i, forms := range arrform {
		arrform[i] = stringchecker(forms)
	}
	res, err := db.Prepare("INSERT INTO games VALUES(?, ?, ?, ?, ?)")
	check(err)
	_, err = res.Exec(arrform[0], arrform[1], arrform[2], arrform[3], arrform[4])
	check(err)
	gethome(rw, r, nil) //Fix later
}

//This should never be run with the migration script working
func connectdb() *sql.DB {
	if db == nil {
		ini, err := ioutil.ReadFile("./res/pass.txt")
		inistr := string(ini)
		check(err)
		db, err = sql.Open("mysql", inistr) //PASSWORD!
		check(err)
		return db
	}
	fmt.Println("Problem: Migration Script Not Run")
	return db
}
func getmessage(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
func postmessage(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func main() {
	gettwilauth()
	router := httprouter.New()
	router.GET("/", getroot)
	router.GET("/home", gethome)
	router.POST("/home", posthome)
	router.GET("/message", getmessage)
	router.POST("/message", postmessage)
	newMessage()
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}
func gettwilauth() {
	file, err := ioutil.ReadFile("./res/twilauth.txt")
	check(err)
	sid, err := fmt.Scanln(file)
	check(err)
	token, err := fmt.Scanln(file)
	check(err)
	twilAuth.sid = string(sid)
	twilAuth.token = string(token)
}

func newMessage() {
	twilio := gotwilio.NewTwilioClient(twilAuth.sid, twilAuth.token)
	response, _, _ := twilio.GetSMS("https://chat.twilio.com/v2/Services")
	fmt.Println(response)
}
