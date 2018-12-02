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

	bits := 14
	pxcm := 0.0099
	tmin := 0.5

	Iraw := Gray16ToFloat(imgOrig)
	theta := proc.AxisAngle
	offset := proc.AxisOffset
	if proc.AxisMethod == "autoDetect" {
		theta, offset = FindCoreAxis(Iraw)
	}
	tmodel := Tmodel(theta, offset, pxcm, proc.CoreType, proc.CoreDiameter, proc.SrcHeight, proc.CoreHeight)
	murhot := MuRhoT(Iraw, bits)
	murhotref := Compensation(murhot, tmodel, tmin)
	low, peak, high := MuRhoTbounds(proc.Low, proc.Mid, proc.High, bits)
	Iproc := ContrastAdjustment(murhotref, low, peak, high)
	Iout := AddScaleBars(Iproc, pxcm, proc.CoreDiameter, proc.SrcHeight, proc.CoreHeight, proc.Motion)
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
