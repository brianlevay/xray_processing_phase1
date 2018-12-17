package main

import (
	"math"
)

func (proc *ImgProcessor) CreateScaleBars() {
	cmPxHeight := cmCoreToPx(proc, 1.0)
	cmPxWidth := cmCoreToPx(proc, proc.Cfg.ScaleWidth)
	roiPxHeight := cmCoreToPx(proc, proc.Cfg.Motion)
	roiPxWidth := cmCoreToPx(proc, proc.Cfg.RoiWidth)
	roiIstart := int(float64(proc.Cfg.HeightPxDet-roiPxHeight) / 2.0)
	iBorder := []int{0, (proc.Cfg.HeightPxDet - 1)}
	jBorder := []int{0, (proc.Cfg.BorderPx + cmPxWidth + proc.Cfg.BorderPx - 1)}
	iCms := []int{(iBorder[0] + proc.Cfg.BorderPx), (iBorder[1] - proc.Cfg.BorderPx)}
	jCms := []int{(jBorder[0] + proc.Cfg.BorderPx), (jBorder[1] - proc.Cfg.BorderPx)}
	iLn := []int{roiIstart, (roiIstart + roiPxHeight - 1)}
	jLn := []int{(jBorder[1] + 1), (jBorder[1] + roiPxWidth)}

	Iscale := make([][]uint16, proc.Cfg.HeightPxDet)
	for i := 0; i < proc.Cfg.HeightPxDet; i++ {
		Iscale[i] = make([]uint16, proc.Cfg.WidthPxDet)
		for j := 0; j < proc.Cfg.WidthPxDet; j++ {
			Iscale[i][j] = 1
			if isInside(i, j, iBorder, jBorder) {
				Iscale[i][j] = 0
			}
			if isInside(i, j, iCms, jCms) {
				if evenCm(i, cmPxHeight) {
					Iscale[i][j] = 0
				} else {
					Iscale[i][j] = 2
				}
			}
			if isInside(i, j, iLn, jLn) {
				Iscale[i][j] = 0
			}
		}
	}
	proc.Iscale = Iscale
	return
}

func isInside(i int, j int, iBounds []int, jBounds []int) bool {
	if (i >= iBounds[0]) && (i <= iBounds[1]) && (j >= jBounds[0]) && (j <= jBounds[1]) {
		return true
	} else {
		return false
	}
}

func evenCm(i int, cmPxHeight int) bool {
	x := math.Floor(float64(i) / float64(cmPxHeight))
	rem := math.Mod(x, 2.0)
	if rem == 0.0 {
		return true
	} else {
		return false
	}
}
