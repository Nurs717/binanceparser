package handler

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		if r.URL.Path != "/" {
			http.Error(w, "Error 404", http.StatusNotFound)
		}

		if r.URL.Path == "/" {
			http.ServeFile(w, r, "web/index.html")
		}
	}
}
