package main

import (
	fe "fileExplorer"
	"log"
	"net/http"
	_ "net/http/pprof"
	"path/filepath"
)

func main() {
	cwd, errCwd := filepath.Abs("./")
	if errCwd != nil {
		log.Fatal(errCwd)
	}
	contents := fe.NewExplorer(cwd, ".tif")

	cfg, errCfg := readConfigToMap("setup.cfg")
	if errCfg != nil {
		log.Fatal(errCfg)
	}

	port := ":8080"
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	fileExplorerHandler(contents)
	histogramHandler(contents, cfg)
	processingHandler(contents, cfg)

	http.ListenAndServe(port, nil)
}
