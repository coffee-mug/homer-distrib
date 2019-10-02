package fileserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFileServer(t *testing.T) {
	t.Run("GET /echo/$message should return message", func(t *testing.T) {
		message := "Hello"
		request, _ := http.NewRequest(http.MethodGet, "/echo/"+message, nil)
		response := httptest.NewRecorder()

		FileServer(response, request)

		got := response.Body.String()

		if got != message {
			t.Errorf("Error: got %s want %s", got, message)
		}
	})

	t.Run("GET /files/$file.txt should return the content of the file", func(t *testing.T) {
		file := "hello.txt"
		expected := "HELLO DISTRIBUTED"

		request, _ := http.NewRequest(http.MethodGet, "/files/"+file, nil)
		response := httptest.NewRecorder()

		FileServer(response, request)

		got := response.Body.String()

		if got != expected {
			t.Errorf("Error: got %s want %s", got, expected)
		}

	})
}
