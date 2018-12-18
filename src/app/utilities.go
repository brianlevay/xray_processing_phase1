package main

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func getAvailablePort() (net.Listener, error) {
	var listener net.Listener
	var errListen error
	preferredPorts := []string{":8080", ":8081", ":8082", ":6060", ":0"}
	for p := 0; p < len(preferredPorts); p++ {
		listener, errListen = net.Listen("tcp", preferredPorts[p])
		if errListen == nil {
			return listener, nil
		}
	}
	return listener, errors.New("Unable to find available port")
}

func startupDisplay(port int) {
	fmt.Println(`__________________________________________________________________`)
	fmt.Println(` __  __    ____               ___                                 `)
	fmt.Println(` \ \/ /   |  _ \ __ _ _   _  |_ _|_ __ ___   __ _  __ _  ___ _ __ `)
	fmt.Println(`  \  / __ |_|_) / _' | | | |  | || '_ ' _ \ / _' |/ _' |/ _ \ '__|`)
	fmt.Println(`  /  \|__||  _ < (_| | |_| |  | || | | | | | (_| | (_| |  __/ |   `)
	fmt.Println(` /_/\_\   |_| \_\__,_|\__, | |___|_| |_| |_|\__,_|\__, |\___|_|   `)
	fmt.Println(`                      |___/                       |___/           `)
	fmt.Println(`__________________________________________________________________`)
	fmt.Println("\nWelcome to the X-Ray Image Processing Program\n")
	fmt.Println("Please access the program via:")
	fmt.Println("http://localhost:" + strconv.Itoa(port) + "/")
	fmt.Println("\nType 'Ctrl+C' or close this window to terminate the program.\n")
	fmt.Println("PROGRAM LOG:")
	return
}

func stringToSlice(valString string) []string {
	replacer := strings.NewReplacer("[", "", "]", "", "\"", "")
	cleaned := replacer.Replace(valString)
	values := strings.Split(cleaned, ",")
	if (len(values) == 1) && (strings.Compare(values[0], "") == 0) {
		return []string{}
	}
	return values
}

func cmCoreToPx(proc *ImgProcessor, cm float64) int {
	return int((cm * proc.ProjMult) / proc.Cfg.CmPerPx)
}

func pxToCmCore(proc *ImgProcessor, px int) float64 {
	return (float64(px) * proc.Cfg.CmPerPx) / proc.ProjMult
}
