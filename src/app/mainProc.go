package main

import (
	fe "fileExplorer"
	"log"
	"math"
	"path"
	"strconv"
	"strings"
)

func ProcessTiffs(contents *fe.FileContents, proc *ImgProcessor, batchN float64) {
	nfiles := len(contents.Selected)
	for i := 0; i < nfiles; i++ {
		proc.ProcessImage(contents.Root, contents.Selected[i])
		if (i != 0) && (math.Mod(float64(i), batchN) == 0) {
			log.Println(strconv.Itoa(i) + " files completed")
		}
	}
	return
}

func (proc *ImgProcessor) ProcessImage(root string, filename string) {
	imgOrig, errOpen := fe.OpenTiff(root, filename)
	if errOpen != nil {
		log.Println(errOpen)
		return
	}
	Iraw, errType := Gray16ToUint16(imgOrig)
	if errType != nil {
		log.Println(errType)
		return
	}
	// This is a stop-gap measure to prevent out-of-bounds errors if the user attempts to process an image of a different size
	pxHeight := len(Iraw)
	pxWidth := len(Iraw[0])
	if (pxHeight != proc.Cfg.HeightPxDet) || (pxWidth != proc.Cfg.WidthPxDet) {
		log.Println("Image dimensions don't match those specified in the configuration.", filename, "will not be processed.")
		return
	}

	theta := proc.AxisAngle
	offset := proc.AxisOffset
	if proc.AxisMethod == "autoDetect" {
		theta, offset = FindCoreAxis(proc, Iraw)
	}
	tmodel := CalculateTModel(proc, theta, offset)
	Iout := PrimaryCalcs(proc, Iraw, tmodel)
	imgOut := Uint16ToGray16(Iout)

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
	return
}
