package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var index = template.Must(template.ParseFiles("index.min.html"))

func init() {

	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.Handle("/js/", http.StripPrefix("/js/", fs))
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	err := index.Execute(w, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
