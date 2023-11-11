package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"portfolia/lib"
	"portfolia/routes"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tmpl.Execute(w, nil)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.IndexHandler)

	loggedMux := lib.NewLogger(mux)
	fmt.Println("Server started on port " + port)
	http.ListenAndServe(":"+port, loggedMux)
}
