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
			proc := new(img.ImgProcessor)
			errJSON := json.Unmarshal([]byte(settingsS[0]), proc)
			errorResponse(errJSON, &w)
			if (proc.SrcHeight == 0.0) || (proc.CoreDiameter == 0.0) {
				invalidValueResponse("Source and/or Core Diameter are invalid", &w)
			}
			if proc.SrcHeight < (proc.CoreHeight + proc.CoreDiameter) {
				invalidValueResponse("Invalid geometry between source and core", &w)
			}
			errSub := fe.CreateSubfolder(contents.Root, proc.FolderName)
			errorResponse(errSub, &w)
			log.Println("Started processing " + strconv.Itoa(nImages) + " images...")
			img.ProcessTiffs(contents, proc)
			log.Println("Finished processing images.")
		}
		log.Println("No images selected.")
		w.Write([]byte(""))
		return
	})
}
