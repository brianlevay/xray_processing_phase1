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

	Iraw := Gray16ToFloat(imgOrig)
	theta := proc.AxisAngle
	offset := proc.AxisOffset
	//if proc.AxisMethod == "autoDetect" {
	//	theta, offset = FindCoreAxis(Iraw)
	//}
	tmodel := Tmodel(proc, Iraw, theta, offset)
	murhot := MuRhoT(proc, Iraw)
	murhotref := Compensation(proc, murhot, tmodel)
	Iproc := ContrastAdjustment(proc, murhotref)
	Iout := Iproc //////////////////////////////
	//Iout := AddScaleBars(proc, Iproc)
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
