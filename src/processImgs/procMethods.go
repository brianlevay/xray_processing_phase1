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

	raw := Gray16ToFloat(imgOrig)
	theta := proc.AxisAngle
	offset := proc.AxisOffset
	if proc.AxisMethod == "autoDetect" {
		theta, offset = FindCoreAxis(raw)
	}
	tmodel := CalculateTmodel(theta, offset, pxcm, proc.CoreType, proc.CoreDiameter, proc.SrcHeight, proc.CoreHeight)
	murhot := CalculateMuRhoT(raw, bits)
	compensated := Compensation(murhot, tmodel, tmin)
	low, peak, high := CalculateMuRhoTbounds(proc.Low, proc.Mid, proc.High, bits)
	processed := AdjustedAndScaled(compensated, low, peak, high)
	imgNew := FloatToGray16(processed)

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
