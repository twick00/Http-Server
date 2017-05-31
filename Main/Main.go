package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/httprouter"
)

var db *sql.DB

// type dataContainer struct {
// 	Title    string
// 	Games    []string
// 	Comments []string
// 	Yesman   string
// }

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

// func decodejson() {
// 	jlist, _ := ioutil.ReadFile("./More/list.json")
// 	var marshjlist = []byte(jlist)
// 	err := json.Unmarshal(marshjlist, &data)
// 	check(err)
// }
// func encodejson() {
// 	marshjlist, err := json.Marshal(data)
// 	check(err)
// 	fmt.Println(marshjlist)
// 	ioutil.WriteFile("./More/list.json", marshjlist, os.ModeTemporary)
// 	//Note: What is mode temp vs perm?
// }

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
	//decodejson()
	t, _ := template.ParseFiles("./More/content.html")
	s1 := t.Lookup("content.html")
	s1.ExecuteTemplate(rw, "newcomment", data)
}
func getroot(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//decodejson()
	var data Gamesstruct
	data.Games = gettabledata()
	log.Print(data.Games[0].Title)
	t, err := template.ParseFiles("./More/content.html")
	check(err)
	s1 := t.Lookup("content.html")
	s1.ExecuteTemplate(rw, "title", data)
}
func getgames(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//decodejson()
	t, _ := template.ParseFiles("./More/content.html")
	s1 := t.Lookup("content.html")
	fmt.Println(data.Games)
	s1.ExecuteTemplate(rw, "games", data)
}
func getcomments(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//decodejson()
	t, _ := template.ParseFiles("./More/content.html")
	s1 := t.Lookup("content.html")
	fmt.Println(data.Comments)
	s1.ExecuteTemplate(rw, "comments", data)
}

func posthome(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//decodejson()
	comment := r.FormValue("comment")
	if comment != "" { //Can create empty strings when pressing the button
		data.Comments = append(data.Comments, comment)
	}
	fmt.Println(data.Comments)

	//encodejson()
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
	var err error
	if db == nil {
		db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/data") //PASSWORD!
		check(err)
		return db
	}
	return db
}

func main() {
	connectdb()
	router := httprouter.New()
	router.GET("/", getroot)
	router.GET("/home", gethome)
	router.GET("/games", getgames)
	router.GET("/games/comments", getcomments)
	router.POST("/home", posthome)

	/*http.HandleFunc("/jump", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("webpage").Parse()
		data.Yesman = "Jumped"
		t.Execute(rw, data)
	})*/
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}
