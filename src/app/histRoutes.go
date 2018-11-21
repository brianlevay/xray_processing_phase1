package main

import (
	"encoding/base64"
	"errors"
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

		selectedS, selPresent := r.Form["Selected"]
		if selPresent == false {
			errSelected := errors.New("Selected variable is not present")
			errorResponse(errSelected, &w)
		}
		contents.Selected = stringToSlice(selectedS[0])
		nImages := len(contents.Selected)

		if nImages > 0 {
			bits, errBits := checkAndConvertToInt("Bits", r.Form)
			errorResponse(errBits, &w)
			nbins, errNbins := checkAndConvertToInt("Nbins", r.Form)
			errorResponse(errNbins, &w)
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
