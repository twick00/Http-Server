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
	"github.com/twick00/go_nuts/Main/res"
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

func check(e error) { //Simple error passing
	if e != nil {
		panic(e)
	}
}

//Game struct contains Title(s), Year(int), Genre(s), Barcode(int) and ID(int)
type Game struct {
	Title   string
	Year    int
	Genre   string
	Barcode int
	ID      int
}

//Gamesstruct Contains []Game
type Gamesstruct struct {
	Games []Game
}

func gethome(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles("./res/content.html")
	check(err)
	s1 := t.Lookup("content.html")
	s1.ExecuteTemplate(rw, "postgamedata", data)
}
func getroot(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var data Gamesstruct
	data.Games = gettabledata()
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
func gettabledata() []Game {
	games := []Game{} //All the data
	game := Game{
		Title:   "",
		Year:    0,
		Genre:   "",
		Barcode: 0,
		ID:      0,
	} //Just one set of data
	rows, err := db.Query("SELECT * FROM test.games") //
	check(err)
	for rows.Next() {
		err := rows.Scan(&game.Title, &game.Year, &game.Genre, &game.Barcode, &game.ID)
		games = append(games, game)
		check(err)
	}
	return games
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

func main() {
	if db == nil {
		db = dbconnect.Command()
		db.Exec("USE test")
	} else {
		connectdb()
	}
	router := httprouter.New()
	router.GET("/", getroot)
	router.GET("/home", gethome)
	router.POST("/home", posthome)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}
