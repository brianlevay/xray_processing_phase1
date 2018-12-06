package main

import (
	"encoding/json"
	fe "fileExplorer"
	"log"
	"net/http"
	img "processImgs"
	"strconv"
)

func processingHandler(contents *fe.FileContents) {
	http.HandleFunc("/processing", func(w http.ResponseWriter, r *http.Request) {
		errP := r.ParseForm()
		errorResponse(errP, &w)
		selectedS, selectedPres := r.Form["Selected"]
		absenceResponse(selectedPres, "Selected", &w)
		contents.Selected = stringToSlice(selectedS[0])
		nImages := len(contents.Selected)
		if nImages > 0 {
			settingsS, settingsPres := r.Form["Settings"]
			absenceResponse(settingsPres, "Settings", &w)
			imgProcessor := new(img.ImgProcessor)
			errJSON := json.Unmarshal([]byte(settingsS[0]), imgProcessor)
			errorResponse(errJSON, &w)
			errSub := fe.CreateSubfolder(contents.Root, imgProcessor.FolderName)
			errorResponse(errSub, &w)
			log.Println("Started processing " + strconv.Itoa(nImages) + " images...")
			img.ProcessTiffs(contents, imgProcessor)
			log.Println("Finished processing images.")
		} else {
			log.Println("No images selected.")
		}
		w.Write([]byte(""))
		return
	})
}
