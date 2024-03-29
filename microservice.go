package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/idiotleon/TutorialCloudNativeGo/api"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)
	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.HandleFunc("/api/books/", api.BookHandleFunc)
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello Cloud Native Go")
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprint(w, message)
}
