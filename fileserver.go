package fileserver

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func FileServer(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/echo/") {
		message := r.URL.Path[len("/echo/"):]
		fmt.Fprintf(w, message)
	}

	if strings.HasPrefix(r.URL.Path, "/files/") {
		if r.Method == "POST" {
			buf := bytes.Buffer{}

			buf.ReadFrom(r.Body)

			// TODO: hardcoded name for now, to update later
			// with more tests
			ioutil.WriteFile("test.txt", buf.Bytes(), 0644)

			fmt.Fprintf(w, buf.String())
		} else {
			filename := r.URL.Path[len("/files/"):]

			http.ServeFile(w, r, filename)
		}
	}
}
