package fileExplorer

import (
	"io/ioutil"
)

func (contents *FileContents) UpdateDir(rootDir string) {
	contents.Root = rootDir
	contents.DirNames = []string{}
	contents.FileNames = []string{}
	contents.Selected = []string{}
	contents.Error = nil

	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		contents.Error = err
		return
	}

	for _, file := range files {
		if file.IsDir() == true {
			contents.DirNames = append(contents.DirNames, file.Name())
		} else {
			contents.FileNames = append(contents.FileNames, file.Name())
		}
	}
	return
}

func NewExplorer(rootDir string) *FileContents {
	contents := new(FileContents)
	contents.UpdateDir(rootDir)
	return contents
}
