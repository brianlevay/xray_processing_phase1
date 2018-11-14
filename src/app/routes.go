package main

import (
	"encoding/json"
	fe "fileExplorer"
	"net/http"
)

func fileExplorerHandler(contents *fe.FileContents) {
	http.HandleFunc("/filesystem", func(w http.ResponseWriter, r *http.Request) {
		var resp []byte
		errP := r.ParseForm()
		if errP != nil {
			w.Write(resp)
			return
		}
		newRoot := r.Form["Root"][0]
		if newRoot != contents.Root {
			contents.UpdateDir(newRoot)
		}
		dataJSON, errJSON := json.Marshal(contents)
		if errJSON != nil {
			w.Write(resp)
			return
		}
		w.Write(dataJSON)
		return
	})
}
