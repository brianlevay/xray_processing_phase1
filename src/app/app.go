package main

import (
	fe "fileExplorer"
	"log"
	"net/http"
	_ "net/http/pprof"
	"path/filepath"
)

func main() {
	cwd, err := filepath.Abs("./")
	if err != nil {
		log.Fatal(err)
	}
	contents := fe.NewExplorer(cwd, ".tif")

	port := ":8080"
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	fileExplorerHandler(contents)
	histogramHandler(contents)
	processingHandler(contents)
	http.ListenAndServe(port, nil)
}
