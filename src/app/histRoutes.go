package main

import (
	"encoding/base64"
	"encoding/json"
	fe "fileExplorer"
	hist "histogram"
	"log"
	"net/http"
)

func histogramHandler(contents *fe.FileContents, cfg map[string]float64) {
	http.HandleFunc("/histogram", func(w http.ResponseWriter, r *http.Request) {
		errP := r.ParseForm()
		if errP != nil {
			log.Println(errP)
			w.Write([]byte(""))
			return
		}
		selectedS, selectedPres := r.Form["Selected"]
		if selectedPres == false {
			log.Println("Selected not present")
			w.Write([]byte(""))
			return
		}
		contents.Selected = stringToSlice(selectedS[0])
		nImages := len(contents.Selected)
		if nImages == 0 {
			log.Println("No images selected.")
			w.Write([]byte(""))
			return
		}
		styleStr, stylePres := r.Form["Style"]
		if stylePres == false {
			log.Println("Style not present")
			w.Write([]byte(""))
			return
		}

		// Create histogram set, fill fields with JSON //
		hset := new(hist.HistogramSet)
		errJSON := json.Unmarshal([]byte(styleStr[0]), hset)
		if errJSON != nil {
			log.Println(errJSON)
			w.Write([]byte(""))
			return
		}
		// Read values in from configuration //
		errInit := hset.Initialize(cfg)
		if errInit != nil {
			log.Println(errInit)
			w.Write([]byte(""))
			return
		}

		// Process files //
		log.Println("Started generating histogram...")
		hist.ImageHistogram(contents, hset)
		sEnc := base64.StdEncoding.EncodeToString(hset.Image)
		log.Println("Sending histogram...")
		w.Write([]byte(sEnc))
		return
	})
}
