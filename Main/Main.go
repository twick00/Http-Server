package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

var o sync.Once
var db *sql.DB
var data = struct {
	Title    string
	Games    []string
	Comments []string
	Yesman   string
}{}

func check(e error) { //Simple error passing
	if e != nil {
		panic(e)
	}
}

//Game is game
type Game struct {
	Title string
	Year  int
	Genre string
}

//Gamesstruct Contains []Game
type Gamesstruct struct {
	Games []Game
}

func gethome(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("./More/content.html")
	s1 := t.Lookup("content.html")
	s1.ExecuteTemplate(rw, "postgamedata", data)
}
func getroot(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var data Gamesstruct
	data.Games = gettabledata()
	t, err := template.ParseFiles("./More/content.html")
	check(err)
	s1 := t.Lookup("content.html")
	s1.ExecuteTemplate(rw, "title", data)
}
func getgames(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("./More/content.html")
	s1 := t.Lookup("content.html")
	fmt.Println(data.Games)
	s1.ExecuteTemplate(rw, "games", data)
}
func getcomments(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("./More/content.html")
	s1 := t.Lookup("content.html")
	fmt.Println(data.Comments)
	s1.ExecuteTemplate(rw, "comments", data)
}

func posthome(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	formTitle := r.FormValue("title")
	formYear := r.FormValue("year")
	formGenre := r.FormValue("genre")
	res, err := db.Prepare("INSERT INTO games VALUES(?, ?, ?)")
	check(err)
	_, err = res.Exec(formTitle, formYear, formGenre)
	check(err)
	gethome(rw, r, nil) //Fix later
}
func gettabledata() []Game {
	games := []Game{} //All the data
	game := Game{
		Title: "",
		Year:  0,
		Genre: "",
	} //Just one set of data
	rows, err := db.Query("SELECT * FROM games") //
	check(err)
	for rows.Next() {
		err := rows.Scan(&game.Title, &game.Year, &game.Genre)
		games = append(games, game)
		check(err)
	}
	return games
}
func connectdb() *sql.DB {
	if db == nil {
		ini, err := ioutil.ReadFile("./More/pass.txt")
		inistr := string(ini)
		check(err)
		db, err = sql.Open("mysql", inistr) //PASSWORD!
		check(err)
		return db
	}
	return db
}

func main() {
	//o.Do(checksetsql)
	connectdb()
	router := httprouter.New()
	router.GET("/", getroot)
	router.GET("/home", gethome)
	router.GET("/games", getgames)
	router.GET("/games/comments", getcomments)
	router.POST("/home", posthome)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}

// func checksetsql() {
// 	migratescript, err := ioutil.ReadFile("./More/migrate.sql")
// 	check(err)
// 	fmt.Println(migratescript)
// 	s := string(migratescript[:])
// 	fmt.Println(s)
// 	_, err = db.Exec(s)
// 	check(err)

// }
