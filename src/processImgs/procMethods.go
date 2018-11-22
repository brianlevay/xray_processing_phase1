package processImgs

import (
	fe "fileExplorer"
	"log"
	"path"
	"strings"
	"sync"
)

func (proc *ImgProcessor) ProcessImage(root string, filename string, wg *sync.WaitGroup) {
	imgOrig, errOpen := fe.OpenTiff(root, filename)
	if errOpen != nil {
		log.Println(errOpen)
		wg.Done()
		return
	}

	raw := Gray16ToFloat(imgOrig)
	imgNew := FloatToGray16(raw)

	rootOut := root
	if proc.FolderName != "" {
		rootOut = path.Join(root, proc.FolderName)
	}
	filenamePts := strings.Split(filename, ".")
	filenameOut := filenamePts[0] + "_processed.tif"
	if proc.FileAppend != "" {
		filenameOut = filenamePts[0] + proc.FileAppend + ".tif"
	}
	errSave := fe.SaveTiff(imgNew, rootOut, filenameOut)
	if errSave != nil {
		log.Println(errSave)
	}
	wg.Done()
	return
}
