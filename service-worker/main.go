package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func getPort() string {
	port := flag.String("port", "8080", "Port to run the server on")
	flag.Parse()
	if *port != "" {
		return ":" + *port
	}
	return ":8000"
}

func main() {
	fmt.Println("Welcome to Go App")
	router := mux.NewRouter()
	http.Handle("/templates/", http.StripPrefix("templates", http.FileServer(http.Dir("./templates"))))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	router.HandleFunc("/", home).Methods("GET")
	http.ListenAndServe(getPort(), router)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err == nil {
		tmpl.Execute(w, nil)
	}
}
