package fileserver

import (
	"fmt"
	"net/http"
)

func FileServer(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path[len("/echo/"):]
	fmt.Fprintf(w, message)
}
