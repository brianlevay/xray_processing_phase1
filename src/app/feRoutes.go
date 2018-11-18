package main

import (
	"encoding/json"
	fe "fileExplorer"
	"net/http"
	"path"
)

func fileExplorerHandler(contents *fe.FileContents) {
	http.HandleFunc("/filesystem", func(w http.ResponseWriter, r *http.Request) {
		errP := r.ParseForm()
		errorResponse(errP, &w)

		var rootPath string
		switch newRoot := r.Form["Root"][0]; newRoot {
		case ".":
			rootPath = contents.Root
		case "..":
			rootPath, _ = path.Split(contents.Root)
			if rootPath == "" {
				rootPath = contents.Root
			} else {
				rootPath = rootPath[0 : len(rootPath)-1]
			}
		default:
			rootPath = path.Join(contents.Root, newRoot)
		}
		contents.UpdateDir(rootPath)

		dataJSON, errJSON := json.Marshal(contents)
		errorResponse(errJSON, &w)
		w.Write(dataJSON)
		return
	})
}
