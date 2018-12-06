package processImgs

import (
	fe "fileExplorer"
	"log"
	"path"
	"strings"
	"sync"
)

func ProcessTiffs(contents *fe.FileContents, proc *ImgProcessor) {
	var wg sync.WaitGroup
	proc.Initialize()
	proc.CreateScaleBars()
	nfiles := len(contents.Selected)
	for i := 0; i < nfiles; i++ {
		wg.Add(1)
		go proc.ProcessImage(contents.Root, contents.Selected[i], &wg)
	}
	wg.Wait()
	return
}

func (proc *ImgProcessor) ProcessImage(root string, filename string, wg *sync.WaitGroup) {
	imgOrig, errOpen := fe.OpenTiff(root, filename)
	if errOpen != nil {
		log.Println(errOpen)
		wg.Done()
		return
	}

	Iraw := Gray16ToFloat(imgOrig)
	theta := proc.AxisAngle
	offset := proc.AxisOffset
	if proc.AxisMethod == "autoDetect" {
		theta, offset = FindCoreAxis(proc, Iraw)
	}
	tmodel := TModel(proc, theta, offset)
	Iout := ProcessByPixel(proc, Iraw, tmodel)
	imgOut := FloatToGray16(Iout)

	rootOut := root
	if proc.FolderName != "" {
		rootOut = path.Join(root, proc.FolderName)
	}
	filenamePts := strings.Split(filename, ".")
	filenameOut := filenamePts[0] + "_processed.tif"
	if proc.FileAppend != "" {
		filenameOut = filenamePts[0] + proc.FileAppend + ".tif"
	}
	errSave := fe.SaveTiff(imgOut, rootOut, filenameOut)
	if errSave != nil {
		log.Println(errSave)
	}
	wg.Done()
	return
}
