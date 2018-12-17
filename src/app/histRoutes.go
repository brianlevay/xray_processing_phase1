package main

import (
	"encoding/base64"
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

		// Create histogram set, read values in from configuration //
		hset := new(hist.HistogramSet)
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
