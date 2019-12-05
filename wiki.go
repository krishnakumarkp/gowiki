package main

import (
	"gowiki2/controller"
	"gowiki2/filestore"
	"log"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view|static)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	controller := controller.WikiController{
		Store: filestore.NewFileStore(),
	}

	http.HandleFunc("/view/", makeHandler(controller.View))
	http.HandleFunc("/edit/", makeHandler(controller.Edit))
	http.HandleFunc("/save/", makeHandler(controller.Save))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
