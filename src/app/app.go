package main

import (
	fe "fileExplorer"
	"net/http"
)

func main() {
	contents := fe.NewExplorer(".")

	port := ":8080"
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	fileExplorerHandler(contents)

	http.ListenAndServe(port, nil)
}
