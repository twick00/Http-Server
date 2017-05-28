package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/httprouter"
)

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

func decodejson() {
	jlist, _ := ioutil.ReadFile("./More/list.json")
	var marshjlist = []byte(jlist)
	err := json.Unmarshal(marshjlist, &data)
	check(err)
}
func encodejson() {
	marshjlist, err := json.Marshal(data)
	check(err)
	fmt.Println(marshjlist)
	ioutil.WriteFile("./More/list.json", marshjlist, os.ModeTemporary)
	//Note: What is mode temp vs perm?
}

func gethome(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decodejson()
	t, _ := template.ParseFiles("./More/content.html")
	s1 := t.Lookup("content.html")
	s1.ExecuteTemplate(rw, "newcomment", data)
}
func getroot(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decodejson()
	t, _ := template.ParseFiles("./More/content.html")
	s1 := t.Lookup("content.html")
	s1.ExecuteTemplate(rw, "title", data)

}
func getgames(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decodejson()
	t, _ := template.ParseFiles("./More/content.html")
	s1 := t.Lookup("content.html")
	fmt.Println(data.Games)
	s1.ExecuteTemplate(rw, "games", data)
}
func getcomments(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decodejson()
	t, _ := template.ParseFiles("./More/content.html")
	s1 := t.Lookup("content.html")
	fmt.Println(data.Comments)
	s1.ExecuteTemplate(rw, "comments", data)
}

func posthome(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decodejson()
	comment := r.FormValue("comment")
	if comment != "" { //Can create empty strings when pressing the button
		data.Comments = append(data.Comments, comment)
	}
	fmt.Println(data.Comments)
	encodejson()
	gethome(rw, r, nil) //Fix later
}

func main() {
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
