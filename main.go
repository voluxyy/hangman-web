package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("templates/Hangman-Homev2.html"))

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.Execute(w, nil)
		if err != nil {
			return
		}

		err = r.ParseForm()
		if err != nil {
			return
		}
	})

	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {

		name := r.FormValue("test")
		println("name : ", name)

		err := r.ParseForm()
		if err != nil {
			return
		}

		//fmt.Println(r.Form)

		//test := r.FormValue("test")
	})

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fmt.Println("[http://localhost:3000] Server run on port 3000")
	http.ListenAndServe(":3000", mux)

}
