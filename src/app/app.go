package main

import (
	"net/http"
)

func main() {
	port := ":8080"
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.ListenAndServe(port, nil)
}
