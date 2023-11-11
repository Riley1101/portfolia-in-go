package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"portfolia/lib"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/parts/base.tmpl",
		"templates/parts/header.tmpl",
		"templates/home.tmpl",
	}

	articles := lib.LoadArticles()

	templates := template.Must(template.ParseFiles(files...))

	fmt.Println(articles)
	if err := templates.ExecuteTemplate(w, "base", articles.Articles); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
