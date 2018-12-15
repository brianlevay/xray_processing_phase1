package main

import (
	"encoding/base64"
	"encoding/json"
	fe "fileExplorer"
	hist "histogram"
	"log"
	"net/http"
)

func histogramHandler(contents *fe.FileContents) {
	http.HandleFunc("/histogram", func(w http.ResponseWriter, r *http.Request) {
		errP := r.ParseForm()
		errorResponse(errP, &w)
		selectedS, selectedPres := r.Form["Selected"]
		absenceResponse(selectedPres, "Selected", &w)
		contents.Selected = stringToSlice(selectedS[0])
		nImages := len(contents.Selected)
		if nImages > 0 {
			sizeStr, sizePres := r.Form["Style"]
			absenceResponse(sizePres, "Style", &w)
			hset := new(hist.HistogramSet)
			errJSON := json.Unmarshal([]byte(sizeStr[0]), hset)
			errorResponse(errJSON, &w)
			log.Println("Started generating histogram...")
			hist.ImageHistogram(contents, hset)
			sEnc := base64.StdEncoding.EncodeToString(hset.Image)
			log.Println("Sending histogram...")
			w.Write([]byte(sEnc))
		}
		log.Println("No images selected.")
		w.Write([]byte(""))
		return
	})
}
