package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/httprouter"
)

var data = struct {
	Title    string
	Items    []string
	Yesman   string
	Comments []string
}{
	Title: "Video Games",
	Items: []string{
		"Dark Souls",
		"Disgaea 5",
		"World of Warcraft",
	},
	Comments: []string{
		"Test",
	},
	Yesman: "",
}

func gethome(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("./More/content.html")
	s1 := t.Lookup("content.html")
	s1.ExecuteTemplate(rw, "newcomment", data)
}
func getroot(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("./More/content.html")
	s1 := t.Lookup("content.html")
	s1.ExecuteTemplate(rw, "content", data)
}
func posthome(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	comment := r.FormValue("comment")
	data.Comments = append(data.Comments, comment)
	log.Printf(comment)
	gethome(rw, r, nil)
}

func main() {
	router := httprouter.New()
	router.GET("/", getroot)

	router.GET("/home", gethome)

	router.POST("/home", posthome)

	/*http.HandleFunc("/jump", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("webpage").Parse()
		data.Yesman = "Jumped"
		t.Execute(rw, data)
	})*/
	log.Fatal(http.ListenAndServe(":8080", router))
}
