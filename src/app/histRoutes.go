package main

import (
	"encoding/json"
	fe "fileExplorer"
	hist "histogram"
	"log"
	"net/http"
	"strings"
)

func histogramHandler(contents *fe.FileContents) {
	http.HandleFunc("/histogram", func(w http.ResponseWriter, r *http.Request) {
		var resp []byte
		errP := r.ParseForm()
		if errP != nil {
			w.Write(resp)
			log.Println(errP)
			return
		}
		contents.Selected = stringToSlice(r.Form["Selected"][0])
		histogram := hist.ImageHistogram(contents, 14, 1000)
		dataJSON, errJSON := json.Marshal(histogram)
		if errJSON != nil {
			w.Write(resp)
			log.Println(errJSON)
			return
		}
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
