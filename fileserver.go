package fileserver

import (
	"fmt"
	"net/http"
	"strings"
)

func FileServer(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/echo/") {
		message := r.URL.Path[len("/echo/"):]
		fmt.Fprintf(w, message)
	}

	if strings.HasPrefix(r.URL.Path, "/files/") {
		filename := r.URL.Path[len("/files/"):]

		http.ServeFile(w, r, filename)
	}
}
