package main

import (
	"encoding/json"
	"errors"
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

		selectedS, selPresent := r.Form["Selected"]
		if selPresent == false {
			errSelected := errors.New("Selected variable is not present")
			errorResponse(errSelected, &w)
		}
		contents.Selected = stringToSlice(selectedS[0])
		nImages := len(contents.Selected)

		settingsS, setPresent := r.Form["Settings"]
		if setPresent == false {
			errSettings := errors.New("Settings are not present")
			errorResponse(errSettings, &w)
		}

		imgProcessor := new(proc.ImgProcessor)
		errJSON := json.Unmarshal([]byte(settingsS[0]), imgProcessor)
		if errJSON != nil {
			errorResponse(errJSON, &w)
		}
		imgProcessor.Initialize()
		errSub := fe.CreateSubfolder(contents.Root, imgProcessor.FolderName)
		if errSub != nil {
			errorResponse(errSub, &w)
		}

		if nImages > 0 {
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
