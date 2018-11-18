package main

import (
	"encoding/json"
	fe "fileExplorer"
	hist "histogram"
	"net/http"
	"strconv"
	"strings"
)

func histogramHandler(contents *fe.FileContents) {
	http.HandleFunc("/histogram", func(w http.ResponseWriter, r *http.Request) {
		errP := r.ParseForm()
		errorResponse(errP, &w)
		contents.Selected = stringToSlice(r.Form["Selected"][0])
		bits, errBits := strconv.Atoi(r.Form["Bits"][0])
		errorResponse(errBits, &w)
		nbins, errNbins := strconv.Atoi(r.Form["Nbins"][0])
		errorResponse(errNbins, &w)
		histogram := hist.ImageHistogram(contents, bits, nbins)
		dataJSON, errJSON := json.Marshal(histogram)
		errorResponse(errJSON, &w)
		w.Write(dataJSON)
		return
	})
}

func stringToSlice(valString string) []string {
	replacer := strings.NewReplacer("[", "", "]", "", "\"", "")
	cleaned := replacer.Replace(valString)
	values := strings.Split(cleaned, ",")
	return values
}
