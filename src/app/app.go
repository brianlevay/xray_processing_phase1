package main

import (
	fe "fileExplorer"
	"fmt"
	"net/http"
)

func main() {
	contents := fe.NewExplorer(".")
	fmt.Println(contents.DirNames)
	fmt.Println(contents.FileNames)

	port := ":8080"
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.ListenAndServe(port, nil)
}
