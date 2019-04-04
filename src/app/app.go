package main

import (
	fe "fileExplorer"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
)

func main() {
	exec, errExec := os.Executable()
	if errExec != nil {
		log.Fatal(errExec)
	}
	cwd, _ := filepath.Split(exec)
	contents := fe.NewExplorer(cwd, ".tif")

	cfg, errCfg := readConfigFromFile(filepath.Join(cwd, "setup.cfg"))
	if errCfg != nil {
		log.Fatal(errCfg)
	}

	fs := http.FileServer(http.Dir(filepath.Join(cwd, "static")))
	http.Handle("/", fs)
	fileExplorerHandler(contents)
	histogramHandler(contents, cfg)
	processingHandler(contents, cfg)
	setupHandler(cfg)

	listener, errListen := getAvailablePort()
	if errListen != nil {
		log.Fatal(errListen)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	startupDisplay(port)
	panic(http.Serve(listener, nil))
}
