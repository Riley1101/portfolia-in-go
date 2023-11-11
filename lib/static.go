package lib

import (
	"net/http"
)

func ServeStatic(mux *http.ServeMux, path string, dir string) {
	mux.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(dir))))
}
