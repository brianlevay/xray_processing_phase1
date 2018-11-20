package main

import (
	"errors"
	fe "fileExplorer"
	"log"
	"net/http"
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

		if nImages > 0 {
			log.Println("Started processing " + strconv.Itoa(nImages) + " images...")
			log.Println("Finished processing images.")
		} else {
			log.Println("No images selected.")
		}
		w.Write([]byte(""))
		return
	})
}
