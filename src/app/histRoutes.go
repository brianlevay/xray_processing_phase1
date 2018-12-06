package main

import (
	"encoding/base64"
	fe "fileExplorer"
	hist "histogram"
	"log"
	"net/http"
	"strings"
)

func histogramHandler(contents *fe.FileContents) {
	http.HandleFunc("/histogram", func(w http.ResponseWriter, r *http.Request) {
		errP := r.ParseForm()
		errorResponse(errP, &w)
		selectedS, selectedPres := r.Form["Selected"]
		absenceResponse(selectedPres, "Selected", &w)
		contents.Selected = stringToSlice(selectedS[0])
		nImages := len(contents.Selected)
		bits := 14
		nbins := 256
		if nImages > 0 {
			width, errWidth := checkAndConvertToInt("Width", r.Form)
			errorResponse(errWidth, &w)
			height, errHeight := checkAndConvertToInt("Height", r.Form)
			errorResponse(errHeight, &w)
			log.Println("Started generating histogram...")
			histogram := hist.ImageHistogram(contents, bits, nbins)
			buffer := hist.DrawHistogram(histogram, width, height)
			sEnc := base64.StdEncoding.EncodeToString(buffer.Bytes())
			log.Println("Sending histogram...")
			w.Write([]byte(sEnc))
		} else {
			log.Println("No images selected.")
			w.Write([]byte(""))
		}
		return
	})
}

func stringToSlice(valString string) []string {
	replacer := strings.NewReplacer("[", "", "]", "", "\"", "")
	cleaned := replacer.Replace(valString)
	values := strings.Split(cleaned, ",")
	if (len(values) == 1) && (strings.Compare(values[0], "") != 0) {
		return []string{}
	}
	return values
}
