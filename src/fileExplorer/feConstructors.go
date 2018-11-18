package fileExplorer

import ()

func NewExplorer(rootDir string, extension string) *FileContents {
	contents := new(FileContents)
	contents.Extension = extension
	contents.UpdateDir(rootDir)
	return contents
}
