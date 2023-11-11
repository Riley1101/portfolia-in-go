package routes

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
}
