package main

import (
	"errors"
	"net"
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
