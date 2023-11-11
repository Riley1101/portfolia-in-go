package main

import (
	"fmt"
	"net/http"
	"os"
	"portfolia/lib"
	"portfolia/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	mux := http.NewServeMux()

	lib.ServeStatic(mux, "/assets/", "assets")
	mux.HandleFunc("/", routes.HomeHandler)
	loggedMux := lib.NewLogger(mux)
	fmt.Println("Server started on port " + port)
	http.ListenAndServe(":"+port, loggedMux)
}
