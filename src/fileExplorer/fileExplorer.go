package fileExplorer

import (
	"io/ioutil"
	"path/filepath"
)

func (contents *FileContents) UpdateDir(rootDir string) {
	var dirNames []string = []string{}
	var fileNames []string = []string{}

	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
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
	contents.Error = nil
	return
}

func NewExplorer(rootDir string, extension string) *FileContents {
	contents := new(FileContents)
	contents.Extension = extension
	contents.UpdateDir(rootDir)
	return contents
}
