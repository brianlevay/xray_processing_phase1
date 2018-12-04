package processImgs

import (
	"math"
)

func AddScaleBars(proc *ImgProcessor, Iproc [][]float64) [][]float64 {
	borderPx := 2
	scaleWidthCm := 0.2
	lineWidthCm := 0.1

	height := len(Iproc)
	width := len(Iproc[0])
	cmPxHeight := pxDet(proc, 1.0)
	cmPxWidth := pxDet(proc, scaleWidthCm)
	lnPxHeight := pxDet(proc, proc.Motion)
	lnPxWidth := pxDet(proc, lineWidthCm)
	lnIstart := int(float64(height-lnPxHeight) / 2.0)

	iBorder := []int{0, (height - 1)}
	jBorder := []int{0, (borderPx + cmPxWidth + borderPx - 1)}
	iCms := []int{(iBorder[0] + borderPx), (iBorder[1] - borderPx)}
	jCms := []int{(jBorder[0] + borderPx), (jBorder[1] - borderPx)}
	iLn := []int{lnIstart, (lnIstart + lnPxHeight - 1)}
	jLn := []int{(jBorder[1] + 1), (jBorder[1] + lnPxWidth)}

	Iout := make([][]float64, height)
	for i := 0; i < height; i++ {
		Iout[i] = make([]float64, width)
		for j := 0; j < width; j++ {
			Iout[i][j] = Iproc[i][j]
			if isInside(i, j, iBorder, jBorder) {
				Iout[i][j] = 0.0
			}
			if isInside(i, j, iCms, jCms) {
				if evenCm(i, cmPxHeight) {
					Iout[i][j] = 0.0
				} else {
					Iout[i][j] = proc.Imax
				}
			}
			if isInside(i, j, iLn, jLn) {
				Iout[i][j] = 0.0
			}
		}
	}
	return Iout
}

func pxDet(proc *ImgProcessor, cmCore float64) int {
	cmDet := (cmCore / (proc.SrcHeight - proc.CoreHeight - (proc.CoreDiameter / 2.0))) * proc.SrcHeight
	pxDet := int(cmDet / proc.Pxcm)
	return pxDet
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
