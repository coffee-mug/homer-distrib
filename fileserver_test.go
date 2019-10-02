package fileserver

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// utils
func assertResponseBody(t *testing.T, response *httptest.ResponseRecorder, expected string) {
	got := response.Body.String()

	if got != expected {
		t.Errorf("Error: got %s want %s", got, expected)
	}
}

func TestFileServer(t *testing.T) {
	t.Run("GET /echo/$message should return message", func(t *testing.T) {
		message := "Hello"
		request, _ := http.NewRequest(http.MethodGet, "/echo/"+message, nil)
		response := httptest.NewRecorder()

		FileServer(response, request)

	})

	t.Run("GET /files/$file.txt should return the content of the file", func(t *testing.T) {
		file := "hello.txt"
		expected := "HELLO DISTRIBUTED"

		request, _ := http.NewRequest(http.MethodGet, "/files/"+file, nil)
		response := httptest.NewRecorder()

		FileServer(response, request)

		assertResponseBody(t, response, expected)
	})

	t.Run("POST /files should upload the file, with filename test.txt, to the server", func(t *testing.T) {
		fileName := "test.txt"
		// TODO: push that test further to mimick command line file arguments
		fileContent, err := ioutil.ReadFile("input_test_file.txt")
		if err != nil {
			log.Fatal("Could not open input_test_file.txt")
		}

		request, _ := http.NewRequest(http.MethodPost, "/files/", ioutil.NopCloser(bytes.NewReader(fileContent)))
		request.Header.Set("Content-Type", "text/plain")

		response := httptest.NewRecorder()

		FileServer(response, request)

		assertResponseBody(t, response, string(fileContent))

		// See if the file has been correctly saved
		request, _ = http.NewRequest(http.MethodGet, "/files/"+fileName, nil)
		response = httptest.NewRecorder()

		FileServer(response, request)

		assertResponseBody(t, response, string(fileContent))
	})
}
