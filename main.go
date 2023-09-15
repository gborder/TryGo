package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "Beetlejuice", Director: "Tim Burton"},
				{Title: "Ratatouille", Director: "Brad Bird, Jan Pinkava"},
				{Title: "City of God", Director: "Fernando Meirelles, Kátia Lund"},
			},
		}
		tmpl.Execute(w, films)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})

	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)
	log.Fatal(http.ListenAndServe(":42350", nil))
}
