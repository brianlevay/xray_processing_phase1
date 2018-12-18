package main

import (
	"encoding/json"
	fe "fileExplorer"
	"log"
	"net/http"
	"strconv"
)

func processingHandler(contents *fe.FileContents, cfg *Configuration) {
	http.HandleFunc("/processing", func(w http.ResponseWriter, r *http.Request) {
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
		// Check for the necessary JSON //
		settingsS, settingsPres := r.Form["Settings"]
		if settingsPres == false {
			log.Println("Settings not present")
			w.Write([]byte(""))
			return
		}

		// Create processor, fill fields with JSON //
		proc := new(ImgProcessor)
		errJSON := json.Unmarshal([]byte(settingsS[0]), proc)
		if errJSON != nil {
			log.Println(errJSON)
			w.Write([]byte(""))
			return
		}
		// Read values in from configuration, then pre-populate additional struct fields and lookup tables //
		errInit := proc.Initialize(cfg)
		if errInit != nil {
			log.Println(errInit)
			w.Write([]byte(""))
			return
		}
		// Create a subfolder for the output files //
		errSub := fe.CreateSubfolder(contents.Root, proc.FolderName)
		if errSub != nil {
			log.Println(errSub)
			w.Write([]byte(""))
			return
		}

		// Process files //
		batchN := 20.0
		log.Println("Started processing " + strconv.Itoa(nImages) + " images...")
		ProcessTiffs(contents, proc, batchN)
		log.Println("Finished processing images.")
		w.Write([]byte(""))
		return
	})
}
