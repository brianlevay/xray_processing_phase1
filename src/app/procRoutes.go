package main

import (
	"encoding/json"
	fe "fileExplorer"
	"log"
	"net/http"
	img "processImgs"
	"strconv"
)

func processingHandler(contents *fe.FileContents, cfg map[string]float64) {
	http.HandleFunc("/processing", func(w http.ResponseWriter, r *http.Request) {
		errP := r.ParseForm()
		errorResponse(errP, &w)
		selectedS, selectedPres := r.Form["Selected"]
		absenceResponse(selectedPres, "Selected", &w)
		contents.Selected = stringToSlice(selectedS[0])
		nImages := len(contents.Selected)
		if nImages == 0 {
			log.Println("No images selected.")
			w.Write([]byte(""))
			return
		}
		// Check for the necessary JSON //
		settingsS, settingsPres := r.Form["Settings"]
		absenceResponse(settingsPres, "Settings", &w)

		// Create processor, fill fields with JSON //
		proc := new(img.ImgProcessor)
		errJSON := json.Unmarshal([]byte(settingsS[0]), proc)
		errorResponse(errJSON, &w)

		// Read values in from configuration, then pre-populate additional struct fields and lookup tables //
		errInit := proc.Initialize(cfg)
		errorResponse(errInit, &w)

		// Create a subfolder for the output files //
		errSub := fe.CreateSubfolder(contents.Root, proc.FolderName)
		errorResponse(errSub, &w)

		// Process files //
		log.Println("Started processing " + strconv.Itoa(nImages) + " images...")
		img.ProcessTiffs(contents, proc)
		log.Println("Finished processing images.")
		w.Write([]byte(""))
		return
	})
}
