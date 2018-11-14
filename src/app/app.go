package main

import (
	fe "fileExplorer"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	cwd, err := filepath.Abs("./")
	if err != nil {
		log.Fatal(err)
	}
	contents := fe.NewExplorer(cwd)

	port := ":8080"
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	fileExplorerHandler(contents)

	http.ListenAndServe(port, nil)
}
