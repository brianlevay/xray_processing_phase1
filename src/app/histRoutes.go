package main

import (
	fe "fileExplorer"
	hist "histogram"
	"net/http"
	"strings"
)

func histogramHandler(contents *fe.FileContents) {
	http.HandleFunc("/histogram", func(w http.ResponseWriter, r *http.Request) {
		var resp []byte
		errP := r.ParseForm()
		if errP != nil {
			w.Write(resp)
			return
		}
		contents.Selected = stringToSlice(r.Form["Selected"][0])
		hist.ImageHistogram(contents, 16, 20)
		w.Write(resp)
		return
	})
}

func stringToSlice(valString string) []string {
	noBrackets := strings.Replace(valString, "[", "", -1)
	noBrackets = strings.Replace(noBrackets, "]", "", -1)
	values := strings.Split(noBrackets, ",")
	return values
}
