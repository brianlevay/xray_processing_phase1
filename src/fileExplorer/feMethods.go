package fileExplorer

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

func (contents *FileContents) UpdateDir(rootDir string) {
	var dirNames []string = []string{}
	var fileNames []string = []string{}
	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		log.Println(err)
		return
	}
	for _, file := range files {
		if file.IsDir() == true {
			dirNames = append(dirNames, file.Name())
		} else {
			if filepath.Ext(file.Name()) == contents.Extension {
				fileNames = append(fileNames, file.Name())
			}
		}
	}
	contents.Root = rootDir
	contents.DirNames = dirNames
	contents.FileNames = fileNames
	contents.Selected = []string{}
	return
}
