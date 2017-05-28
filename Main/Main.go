package main

import (
	"hello/Main/more"
	"html/template"
	"log"
	"net/http"
)

var data = struct {
	Title  string
	Items  []string
	Yesman string
}{
	Title: "Video Games",
	Items: []string{
		"Dark Souls",
		"Disgaea 5",
		"World of Warcraft",
	},
	Yesman: "",
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("webpage").Parse(Templates.Tpl)
		data.Yesman = "Not Jumped"
		t.Execute(rw, data)
	})
	http.HandleFunc("/jump", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("webpage").Parse(Templates.Tpl)
		data.Yesman = "Jumped"
		t.Execute(rw, data)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

}
