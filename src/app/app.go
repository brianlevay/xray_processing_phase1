package main

import (
	fe "fileExplorer"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"path/filepath"
	"strconv"
)

func main() {
	cwd, errCwd := filepath.Abs("./")
	if errCwd != nil {
		log.Fatal(errCwd)
	}
	contents := fe.NewExplorer(cwd, ".tif")

	cfg, errCfg := readConfigFromFile("setup.cfg")
	if errCfg != nil {
		log.Fatal(errCfg)
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	fileExplorerHandler(contents)
	histogramHandler(contents, cfg)
	processingHandler(contents, cfg)

	listener, errListen := getAvailablePort()
	if errListen != nil {
		log.Fatal(errListen)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	log.Println("Program running at:")
	log.Println("http://localhost:" + strconv.Itoa(port) + "/")
	panic(http.Serve(listener, nil))
}
