package main

import (
	"encoding/json"
	fe "fileExplorer"
	"net/http"
	"path"
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
		var rootPath string
		if newRoot == ".." {
			rootPath, _ = path.Split(contents.Root)
			if rootPath == "" {
				rootPath = "."
			} else {
				rootPath = rootPath[0 : len(rootPath)-1]
			}
		} else {
			rootPath = path.Join(contents.Root, newRoot)
		}

		contents.UpdateDir(rootPath)
		dataJSON, errJSON := json.Marshal(contents)
		if errJSON != nil {
			w.Write(resp)
			return
		}
		w.Write(dataJSON)
		return
	})
}
